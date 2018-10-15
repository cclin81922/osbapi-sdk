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
