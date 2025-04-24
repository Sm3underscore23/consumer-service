package service

import (
	"context"
	"event-generator/iternal/model"
)

type OrderSenderService interface {
	SendOrderData(ctx context.Context, orderRq model.OrderRequest) error
}
