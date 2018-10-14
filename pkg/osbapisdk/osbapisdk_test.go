package osbapisdk

import "testing"

func TestEcho(t *testing.T) {
    message := "hi"
    reply, err := Echo(message)
    if err != nil {
        t.Fatal(err)
    }
    if reply != message {
        t.Fatal("reply message error")
    }
}
