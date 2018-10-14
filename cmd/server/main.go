package main

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func root(w http.ResponseWriter, r *http.Request) {
	// log.Println(r.Body)
	io.Copy(w, r.Body)
}

func main() {
	go func() {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		tlsConfig.BuildNameToCertificate()
		transport := &http.Transport{TLSClientConfig: tlsConfig}
		client := &http.Client{Transport: transport}

		for now := range time.Tick(1e9) {
			_ = now
			// resp, _ := http.Post("http://localhost:8080/", "Content-Type: application/x-www-form-urlencoded", strings.NewReader("message"))
			// resp, _ := http.Post("https://localhost.localdomain:8080/", "Content-Type: application/x-www-form-urlencoded", strings.NewReader("message"))
			resp, _ := client.Post("https://localhost.localdomain:8080/", "Content-Type: application/x-www-form-urlencoded", strings.NewReader("message"))
			body, _ := ioutil.ReadAll(resp.Body)
			log.Println(string(body))
		}
	}()

	http.HandleFunc("/", root)
	// http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServeTLS(":8080", "pki/server.cert.pem", "pki/server.key.pem", nil))
}
