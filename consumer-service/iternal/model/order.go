package model

import "time"

type OrderData struct {
	Id     int64  `json:"id"`
	Status string `json:"status"`
}

type Order struct {
	OrderInfo  OrderData
	UpdateTime time.Time `json:"created_at"`
}
