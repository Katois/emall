package main

import (
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"dubbo.apache.org/dubbo-go/v3/server"
	"github.com/katois/emall/api/payment-api/proto/payment"
	"payment-srv/internal/handler"
)

func main() {
	srv, err := server.NewServer(server.WithServerProtocol(protocol.WithPort(20000), protocol.WithTriple()))
	if err != nil {
		panic(err)
	}
	err = payment.RegisterPaymentServiceHandler(srv, handler.NewPaymentTriHandler())
	if err != nil {
		panic(err)
	}
	if err = srv.Serve(); err != nil {
		panic(err)
	}
}
