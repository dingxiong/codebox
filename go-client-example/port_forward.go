package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"io"
	"net/url"
	"path/filepath"
	"strings"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	"k8s.io/client-go/util/homedir"
)

func PortForward() (err error) {
	logger.Info().Msg("Start port forward...")

	podName := "helloworld-v1-77cb56d4b4-9gm9w"
	nameSpace := "sample"
	localPort := "5000"
	remotePort := "5000"
	ports := []string{fmt.Sprintf("%s:%s", localPort, remotePort)}

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String(
			"kubeconfig",
			filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeconfig file",
		)
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return err
	}
	roundTripper, upgrader, err := spdy.RoundTripperFor(config)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/portforward", nameSpace, podName)
	hostIP := strings.TrimLeft(config.Host, "htps:/")
	logger.Info().Msgf("host ip %s", hostIP)
	serverURL := url.URL{Scheme: "https", Path: path, Host: hostIP}

	dialer := spdy.NewDialer(
		upgrader,
		&http.Client{Transport: roundTripper},
		http.MethodPost,
		&serverURL,
	)
	stopChan, readyChan := make(chan struct{}, 1), make(chan struct{}, 1)
	out, errOut := new(bytes.Buffer), new(bytes.Buffer)

	forwarder, err := portforward.New(dialer, ports, stopChan, readyChan, out, errOut)
	if err != nil {
		return
	}

	go func() {
		for range readyChan { // Kubernetes will close this channel when it has something to tell us.
		}
		if len(errOut.String()) != 0 {
			logger.Error().Msgf("%v", errOut)
		} else if len(out.String()) != 0 {
			logger.Info().Msgf("%v", out)
		}
	}()

	go func() {
		if err := forwarder.ForwardPorts(); err != nil { // Locks until stopChan is closed.
			logger.Fatal().Msgf("port forwarder failed: %v", err)
		}
	}()

	<- readyChan
	requestURL := fmt.Sprintf("http://localhost:%s/hello", localPort)
	resp, err := http.Get(requestURL)
	logger.Info().Msgf("%v", resp)
	bodyBytes, err := io.ReadAll(resp.Body)
	logger.Info().Msgf("%s", bodyBytes)

	return
}
