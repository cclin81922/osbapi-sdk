package osbapisdk

import "testing"

func TestEcho(t *testing.T) {
	Setup("../../pki/baseurl", "../../pki/ca.cert.pem", "../../pki/client.key.pem", "../../pki/client.cert.pem")

	message := "hi"
	reply, err := Echo(message)
	if err != nil {
		t.Fatal(err)
	}
	if reply != message {
		t.Fatalf("reply message error | expected %s | got %s", message, reply)
	}
}
