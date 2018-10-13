# Installation

```
go get -u -d github.com/cclin81922/osbapi-sdk/pkg/osbapisdk
```

# Package Usage

```
import "github.com/cclin81922/osbapi-sdk/pkg/osbapisdk"

func demo(message string) {
    reply := osbapisdk.Echo(message)
    fmt.Println(reply)
}
```
