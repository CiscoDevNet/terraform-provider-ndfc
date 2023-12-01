[![Tests](https://github.com/netascode/go-nd/actions/workflows/test.yml/badge.svg)](https://github.com/netascode/go-nd/actions/workflows/test.yml)

# go-nd

`go-nd` is a Go client library for Cisco Nexus Dashboard. It is based on Nathan's excellent [goaci](https://github.com/brightpuddle/goaci) module and features a simple, extensible API and [advanced JSON manipulation](#result-manipulation).

## Getting Started

### Installing

To start using `go-nd`, install Go and `go get`:

`$ go get -u github.com/netascode/go-nd`

### Basic Usage

```go
package main

import "github.com/netascode/go-nd"

func main() {
    client, _ := nd.NewClient("1.1.1.1", "/appcenter/cisco/ndfc/api/v1", "user", "pwd", "", true)

    res, _ := client.Get("/lan-fabric/rest/control/fabrics")
    println(res.Get("0.id").String())
}
```

This will print something like:

```
3
```

#### Result manipulation

`nd.Result` uses GJSON to simplify handling JSON results. See the [GJSON](https://github.com/tidwall/gjson) documentation for more detail.

```go
res, _ := client.Get("/lan-fabric/rest/control/fabrics")

for _, group := range res.Array() {
    println(group.Get("@pretty").String()) // pretty print fabrics
}
```

#### POST data creation

`nd.Body` is a wrapper for [SJSON](https://github.com/tidwall/sjson). SJSON supports a path syntax simplifying JSON creation.

```go
body := nd.Body{}.
    Set("templatename", "test").
    Set("content", "##template properties \nname= test;\ndescription= ;\ntags= ;\nsupportedPlatforms= All;\ntemplateType= POLICY;\ntemplateSubType= VLAN;\ncontentType= TEMPLATE_CLI;##template variables\r\n##\r\n##template content\r\n##")
client.Post("/configtemplate/rest/config/templates/template", body.Str)
```

## Documentation

See the [documentation](https://godoc.org/github.com/netascode/go-nd) for more details.
