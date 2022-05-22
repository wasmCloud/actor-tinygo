module github.com/wasmcloud/actor-tinygo/example

//replace github.com/wasmcloud/actor-tinygo => ../

//replace github.com/wasmcloud/tinygo-msgpack => ../../../../../../wasmcloud-async/tinygo-msgpack

//replace github.com/wasmcloud/interfaces/httpserver/tinygo => ../../../../../../wasmcloud-async/interfaces/httpserver/tinygo

go 1.17

require (
	github.com/wasmcloud/actor-tinygo v0.0.0-20220522040027-3c606438d4bb
	github.com/wasmcloud/interfaces/httpserver/tinygo v0.0.0-20220521004608-a2593fe14af5
)

require github.com/wasmcloud/tinygo-msgpack v0.1.3-0.20220520001432-192ec93ec5c4 // indirect
