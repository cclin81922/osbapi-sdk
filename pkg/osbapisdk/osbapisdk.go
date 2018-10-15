//    Copyright 2018 cclin
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package osbapisdk

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	baseurlFile string
	caFile      string
	keyFile     string
	certFile    string
)

func init() {
	baseurlFile = "/etc/osbapi-svc-credential/baseurl"
	caFile = "/etc/osbapi-svc-credential/ca"
	keyFile = "/etc/osbapi-svc-credential/key"
	certFile = "/etc/osbapi-svc-credential/cert"
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
	resp, err := client.Post(api, "Content-Type: application/x-www-form-urlencoded", strings.NewReader(message))
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
