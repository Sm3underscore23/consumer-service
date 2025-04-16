package repository

import (
	"consumer-service/iternal/model"
	"context"
)

type OrderRepository interface {
	Create(ctx context.Context, order model.OrderData) error
	Get(ctx context.Context, id int64) (model.OrderData, error)
	Update(ctx context.Context, orderData model.OrderData) error
}
