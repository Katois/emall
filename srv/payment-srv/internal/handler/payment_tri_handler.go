package handler

import (
	"context"
	proto "github.com/katois/emall/api/payment-api/proto/payment"
)

type PaymentTriHandler struct {
}

func NewPaymentTriHandler() *PaymentTriHandler {
	return &PaymentTriHandler{}
}

func (h *PaymentTriHandler) Pay(ctx context.Context, request *proto.PayReq) (*proto.PayResp, error) {
	return &proto.PayResp{
		Success: true,
	}, nil
}
