# TinyGo Actor Example - Echo
This example takes an incoming HTTP request and returns a simple HTTP response. This is very similar to the existing echo samples that we have available. Feel free to experiment with the code to see how you can change the HTTP response easily.

## Building
To build this code run the following command:

```
tinygo build -o ./echo.wasm -target wasm -no-debug .
```

Once built, you can execute the `sign.sh` shell script that will use `wash` to sign the module and make it available for running inside a wasmCloud host.