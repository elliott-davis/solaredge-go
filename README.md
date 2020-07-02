# Golang SolarEdge Monitoring API

Go client for the SolarEdge monitoring API

## Usage

```go
package main

import (
    "fmt"
    "os"

    "github.com/elliott-davis/solaredge-go/solaredge"
)

func main() {
    token := os.Getenv("SOLAREDGE_AUTH_TOKEN")
    // You may optionally include your own http client
    client := solaredge.NewClient(nil, token)
    site, err := client.Site.List(&solaredge.ListOptions{Page: 2, PerPage: 1})
    if err != nil {
    	panic(err)
    }
    fmt.Println(site)
}
```
