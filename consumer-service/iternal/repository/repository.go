package repository

import (
	"consumer-service/iternal/model"
	"context"
)

type OrderRepository interface {
	Create(ctx context.Context, order model.Order) error
	Get(ctx context.Context, id int64) (model.Order, error)
	Update(ctx context.Context, orderData model.Order) error
}
