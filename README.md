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
go test github.com/cclin81922/osbapi-sdk/pkg/osbapisdk
```
