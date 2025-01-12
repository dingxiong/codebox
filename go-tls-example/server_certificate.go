package main

import (
	"crypto/tls"
	"fmt"
	"log"
)

func main() {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "webhook-staging.ziphq.com:443", conf)
	if err != nil {
		log.Println("Error in Dial", err)
		return
	}
	defer conn.Close()
	certs := conn.ConnectionState().PeerCertificates
	tls.LoadX509KeyPair()
	for _, cert := range certs {
		fmt.Println("----")
		fmt.Printf("Issuer Name: %s\n", cert.Issuer)
		fmt.Printf("Expiry: %s \n", cert.NotAfter.Format("2006-01-02T15:04:05-0700"))
		fmt.Printf("Common Name: %s \n", cert.Issuer.CommonName)
		fmt.Printf("Ip address: %v\n", cert.IPAddresses)
		fmt.Printf("DNS names: %v\n", cert.DNSNames)
		fmt.Printf("URIs: %v\n", cert.URIs)
	}
	fmt.Println()
	fmt.Printf("server name: %s \n", conn.ConnectionState().ServerName)
}
