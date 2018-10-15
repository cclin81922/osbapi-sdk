# Installation

```
go get -u -d github.com/cclin81922/osbapi-sdk/pkg/osbapisdk
```

# Package Usage

```
import "github.com/cclin81922/osbapi-sdk/pkg/osbapisdk"

func demo(message string) {
    reply, err := osbapisdk.Echo(message)
    fmt.Println(reply)
}
```

# For Developer

Run all tests

```
go get -u github.com/cclin81922/osbapi-baas/cmd/osbapibaas
export PATH=$PATH:~/go/bin
osbapibaas -port=8443

echo "127.0.0.1   localhost.localdomain" >> /etc/hosts
go test github.com/cclin81922/osbapi-sdk/pkg/osbapisdk
```