package model

import "time"

type OrderRequest struct {
	Id     int64  `json:"id"`
	Status string `json:"status"`
}

type Order struct {
	OrderInfo  OrderRequest
	UpdateTime time.Time `json:"created_at"`
}
