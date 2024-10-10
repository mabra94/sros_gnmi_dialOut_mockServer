# SROS Mock Server

This project simulates a Nokia SROS device that sends telemetry data using the gNMI protocol. It focuses specifically on the simulation of SROS gNMI dial out telemetry. 

See this Nokia SROS documentation for more information:
https://documentation.nokia.com/sr/24-7/7x50-shared/system-management/grpc.html?hl=grpc#ai9exgst8x

This is a very experimental version. The only thing it does today is to establish a connection and publish notification following a hardcoded format. 

What it cannot do:
- TLS based transport
- any specific encoding
- ...

## Usage

To run the mock server:

```
go run main.go
```

To specify a custom server address:

```
go run main.go --server=192.168.1.100:57400

```

## Dependencies

- github.com/openconfig/gnmi
- google.golang.org/grpc

