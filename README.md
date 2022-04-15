# wasmCloud TinyGo Actor SDK
This module is for those looking to develop [wasmCloud](https://wasmcloud.dev) actors using [TinyGo](https://tinygo.org). Typically you won't use this module directly, as it will be referenced by one or more code generate modules containing strongly-typed structs and wrappers around function handler registration.

A typical actor might look like this (**NOTE** this is hypothetical at the moment until we have code generation available):

```go
package main

import (
    httpserver "github.com/wasmCloud/interfaces/httpserver/go"
)

func main() {
   httpserver.RegisterHandleRequest(handleRequest)
 }

 func handleRequest(request httpserver.HttpRequest) (httpserver.HttpResponse, error) {
  // process request
  // return response
 }
```