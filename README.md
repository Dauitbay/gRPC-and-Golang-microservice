## gRPC Kitchen & Orders

This project demonstrates a simple gRPC-based microservice setup written in Go.

It consists of:
- **Orders service** — a gRPC server responsible for order management
- **Kitchen service** — an HTTP server that consumes the Orders service via gRPC and renders results as HTML
- **Shared protobuf definitions** with generated Go code

The project is intentionally minimal and avoids frameworks to keep the core ideas clear.

---

## Architecture

Browser
|
v
Kitchen (HTTP) ───── gRPC ─────▶ Orders (gRPC)


- The Kitchen service exposes an HTTP endpoint
- On each request, it:
    - Creates an order via gRPC
    - Fetches orders for a customer
    - Renders them using `html/template`

---

## Requirements

- Go **1.21+**
- `protoc`
- Go protobuf plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Generate gRPC Code
```bash
make gen
```
Generated code will be placed in:
```bash
services/common/genproto/orders
```

Running the Services
#### 1. Start Orders (gRPC) service
```bash
make run-orders
```
Expected output:
```bash
server listening at :9000
```

#### 2. Start Kitchen (HTTP) service

In a separate terminal:
```bash
make run-kitchen
```
Expected output:
```bash 
server listening at :1000
```
#### 3. Open in Browser
```bash
http://localhost:1000
```
You should see a web page displaying a list of orders.

Notes

 - No frameworks are used; only Go standard library and gRPC
 - gRPC client connection is reused across HTTP requests
 - Server-side HTML rendering is done with html/template
 - Errors inside handlers do not crash the server
 
### Common Issues
    connection refused

Make sure the Orders service is running before starting Kitchen.

    exit status 1

Check that:
 - You are not using log.Fatal for normal server execution
 - Ports 9000 and 1000 are not already in use


---

If you want a **shorter README**, a **production-grade one**, or a **Docker-ready version**, say the word.
