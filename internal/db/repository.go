package db

import "context"

type OrderRepository interface {
	GetByID(ctx context.Context)
}
