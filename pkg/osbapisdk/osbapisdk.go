package osbapisdk

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	baseurlFile string
	caFile      string
	keyFile     string
	certFile    string
)

func init() {
	baseurlFile = "/etc/osbapi-svc-credentials/baseurl"
	caFile = "/etc/osbapi-svc-credentials/ca"
	keyFile = "/etc/osbapi-svc-credentials/key"
	certFile = "/etc/osbapi-svc-credentials/cert"
}

// Setup ...
func Setup(baseurl, ca, key, cert string) {
	baseurlFile = baseurl
	caFile = ca
	keyFile = key
	certFile = cert
}

// Echo ...
func Echo(message string) (string, error) {
	baseurlBytes, err := ioutil.ReadFile(baseurlFile)
	if err != nil {
		return "", err
	}
	baseurl := string(baseurlBytes)
	api := fmt.Sprintf("%s/echo", baseurl)

	// Load client cert
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return "", err
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// Send HTTP request
	resp, err := client.Post(api, "", bytes.NewBufferString(message))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	replyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	reply := string(replyBytes)

	return reply, nil
}
