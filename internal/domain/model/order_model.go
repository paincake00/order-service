package model

import "time"

type OrderModel struct {
	OrderUID          string        `json:"order_uid" db:"order_uid"`
	TrackNumber       string        `json:"track_number" db:"track_number"`
	Entry             string        `json:"entry" db:"entry"`
	Delivery          DeliveryModel `json:"delivery" db:"delivery"`
	Payment           PaymentModel  `json:"payment" db:"payment"`
	Items             []ItemModel   `json:"items" db:"items"`
	Locale            string        `json:"locale" db:"locale"`
	CustomerID        string        `json:"customer_id" db:"customer_id"`
	DateCreated       time.Time     `json:"date_created" db:"date_created"`
	InternalSignature string        `json:"internal_signature,omitempty" db:"internal_signature"`
	DeliveryService   string        `json:"delivery_service" db:"delivery_service"`
	ShardKey          string        `json:"shardkey" db:"shardkey"`
	SmID              int           `json:"sm_id" db:"sm_id"`
	OofShard          string        `json:"oof_shard" db:"oof_shard"`
}
