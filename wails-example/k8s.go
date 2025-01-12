package main

import (
	"context"
	"path/filepath"

	"github.com/anthhub/forwarder"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/homedir"
)

func getkubeConfigPath() string {
	home := homedir.HomeDir();
	kubeconfigPath := filepath.Join(home, ".kube", "config")
	return kubeconfigPath
}

func getK8sConfig() *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", getkubeConfigPath())
	if err != nil {
		logger.Fatal().Msgf("Fail to initialize k8s configuration: %+v", err)
	}
	return config
}


func GetSecret(ctx context.Context, namespace, secretName string) (map[string][]byte, error) {
	config := getK8sConfig()
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	secrets, err := clientset.CoreV1().Secrets(namespace).Get(ctx, secretName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return secrets.Data, nil
}

func Portforward(ctx context.Context) *forwarder.Result {
	options := []*forwarder.Option{
		{
			LocalPort:   ES_PRODUCTION_LOCAL_PORT,
			RemotePort:  ES_PRODUCTION_REMOTE_PORT,
			Namespace: "mysql-nginx",
			ServiceName: "es-proxy-prod",
		},
		{
			LocalPort:  ES_STAGING_LOCAL_PORT,
			RemotePort: ES_STAGING_REMOTE_PORT,
			Namespace: "mysql-nginx",
			ServiceName: "es-proxy-staging",
		},
	}

	ret, err := forwarder.WithForwarders(ctx, options, getkubeConfigPath())
	if err != nil {
		panic(err)
	}

	// remember to close the forwarding
	// defer ret.Close()

	// wait forwarding ready
	// the remote and local ports are listed
	ports, err := ret.Ready()
	if err != nil {
		panic(err)
	}
	logger.Info().Msgf("ports: %+v", ports)
	return ret
}
