module github.com/wasmcloud/actor-tinygo/example

replace github.com/wasmcloud/actor-tinygo v0.0.0-20220520002104-dde9cf474f9d => /Users/steve/projects/go/src/github.com/wasmcloud/actor-tinygo

replace github.com/wasmcloud/interfaces/httpserver/tinygo => /Users/steve/projects/wasmcloud-async/interfaces/httpserver/tinygo

go 1.17

require github.com/wasmcloud/actor-tinygo v0.0.0-20220520002104-dde9cf474f9d

require github.com/wasmcloud/tinygo-msgpack v0.1.3-0.20220520001432-192ec93ec5c4 // indirect
