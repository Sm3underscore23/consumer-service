package order

import (
	"consumer-service/iternal/model"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
)

func (s *messageService) processOrder(ctx context.Context, rowOrderData []byte) error {
	var orderData model.OrderData
	err := json.Unmarshal(rowOrderData, &orderData)
	if err != nil {
		return err
	}

	getedOrderData, err := s.repo.Get(ctx, orderData.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			if err := s.repo.Create(ctx, orderData); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	if orderData.Status != getedOrderData.Status {
		if err := s.repo.Update(ctx, orderData); err != nil {
			return nil
		}
	}

	return nil
}
