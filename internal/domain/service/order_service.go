package service

import (
	"context"

	"github.com/paincake00/order-service/internal/db"
	"github.com/paincake00/order-service/internal/domain/model"
)

type OrderService interface {
	GetOrderByID(ctx context.Context, orderUID string) (*model.OrderModel, error)
}

type OrderServiceImpl struct {
	dbRepo db.OrderDBRepository
	// cacheRepo cache.OrderCacheRepository
}

func NewOrderService(orderRepo db.OrderDBRepository) OrderService {
	return &OrderServiceImpl{dbRepo: orderRepo}
}

func (s *OrderServiceImpl) GetOrderByID(ctx context.Context, orderUID string) (*model.OrderModel, error) {
	return s.dbRepo.GetByID(ctx, orderUID)
}
