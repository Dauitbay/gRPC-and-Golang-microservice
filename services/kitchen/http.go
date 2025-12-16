package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"
	
	"github.com/sikozonpc/kitchen/services/common/genproto/orders"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()
	
	conn := NewGRPCClient(":9000")
	defer conn.Close()
	
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)
		
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()
		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerId: 23,
			ProductId:  3290,
			Quantity:   2,
		})
		if err != nil {
			log.Printf("failed to create order error %v", err)
			return
		}
		res, err := c.GetOrders(ctx, &orders.GetOrdersRequest{CustomerId: 23})
		if err != nil {
			log.Printf("client error %v", err)
		}
		t := template.Must(template.New("orders").Parse(orderTemplate))
		if err := t.Execute(w, res.GetOrders()); err != nil {
			log.Printf("template error %v", err)
		}
	})
	log.Printf("server listening at %v", s.addr)
	return http.ListenAndServe(s.addr, router)
}

var orderTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Kitchen Orders</title>

	<style>
		body {
			font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
			background: #f7f7f8;
			padding: 24px;
		}

		h1 {
			margin-bottom: 16px;
		}

		table {
			border-collapse: collapse;
			width: 100%;
			max-width: 720px;
			background: #fff;
			box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
		}

		th, td {
			padding: 10px 14px;
			text-align: left;
		}

		th {
			background: #f0f0f0;
			font-weight: 600;
			border-bottom: 2px solid #ddd;
		}

		tr:nth-child(even) {
			background: #fafafa;
		}

		tr:hover {
			background: #f5f5f5;
		}

		td {
			border-bottom: 1px solid #eee;
		}
	</style>
</head>

<body>
	<h1>Kitchen Orders</h1>

	<table>
		<thead>
			<tr>
				<th>Order ID</th>
				<th>Customer ID</th>
				<th>Quantity</th>
			</tr>
		</thead>

		<tbody>
			{{range .}}
			<tr>
				<td>{{.OrderId}}</td>
				<td>{{.CustomerId}}</td>
				<td>{{.Quantity}}</td>
			</tr>
			{{else}}
			<tr>
				<td colspan="3" style="text-align:center; padding:16px;">
					No orders found
				</td>
			</tr>
			{{end}}
		</tbody>
	</table>
</body>
</html>
`
