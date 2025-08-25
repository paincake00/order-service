package dto

type OrderDTO struct {
	OrderUID          string      `json:"order_uid" validate:"required"`
	TrackNumber       string      `json:"track_number" validate:"required"`
	Entry             string      `json:"entry" validate:"required"`
	Delivery          DeliveryDTO `json:"delivery" validate:"required"`
	Payment           PaymentDTO  `json:"payment" validate:"required"`
	Items             []ItemDTO   `json:"items" validate:"required,min=1,dive"`
	Locale            string      `json:"locale" validate:"required"`
	CustomerID        string      `json:"customer_id" validate:"required"`
	DateCreated       string      `json:"date_created" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	InternalSignature string      `json:"internal_signature,omitempty"`
	DeliveryService   string      `json:"delivery_service" validate:"required"`
	ShardKey          *string     `json:"shardkey" validate:"required"`
	SmID              *int        `json:"sm_id" validate:"required"`
	OofShard          *string     `json:"oof_shard" validate:"required"`
}
