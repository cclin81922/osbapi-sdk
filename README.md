# Installation

```
go get -u -d github.com/cclin81922/osbapi-sdk/pkg/sdk
```

# Package Usage

```
import "github.com/cclin81922/osbapi-sdk/pkg/sdk"

func demo(message string) {
    reply := sdk.Echo(message)
    fmt.Println(reply)
}
```
