package order

import (
	"consumer-service/iternal/model"
	"context"
	"encoding/json"
	"errors"
)

func (s *messageService) processOrder(ctx context.Context, rowOrderData []byte) error {
	var order model.Order
	err := json.Unmarshal(rowOrderData, &order)
	if err != nil {
		return err
	}

	getedOrderData, err := s.repo.Get(ctx, order.OrderInfo.Id)
	if err != nil {
		if !errors.Is(err, model.ErrObjectNotExists) {
			return err
		}
		if err := s.repo.Create(ctx, order); err != nil {
			return err
		}
		return nil
	}

	if order.OrderInfo.Status == getedOrderData.OrderInfo.Status {
		return nil
	}

	if err := s.repo.Update(ctx, order); err != nil {
		return err
	}

	return nil
}
