package service

import (
	"context"

	"github.com/paincake00/order-service/internal/cache"
	"github.com/paincake00/order-service/internal/db"
	"github.com/paincake00/order-service/internal/domain/model"
	"go.uber.org/zap"
)

type OrderService interface {
	GetOrderByID(ctx context.Context, orderUID string) (*model.OrderModel, error)
	RestoreCache(ctx context.Context) error
}

type OrderServiceImpl struct {
	dbRepo    db.OrderDBRepository
	cacheRepo *cache.LRUCache
	logger    *zap.SugaredLogger
}

func NewOrderService(orderRepo db.OrderDBRepository, cacheRepo *cache.LRUCache, logger *zap.SugaredLogger) OrderService {
	return &OrderServiceImpl{
		dbRepo:    orderRepo,
		cacheRepo: cacheRepo,
		logger:    logger,
	}
}

func (s *OrderServiceImpl) GetOrderByID(ctx context.Context, orderUID string) (*model.OrderModel, error) {
	if order, ok := s.cacheRepo.Get(orderUID); ok {
		s.logger.Infof("Cache used for uid: %s", orderUID)

		return order, nil
	}
	order, err := s.dbRepo.GetByID(ctx, orderUID)
	if err != nil {
		return nil, err
	}

	s.cacheRepo.Put(order)

	return order, nil
}

func (s *OrderServiceImpl) RestoreCache(ctx context.Context) error {
	capacity := s.cacheRepo.Capacity
	if capacity == 0 {
		return nil
	}

	orders, err := s.dbRepo.GetNLast(ctx, capacity)
	if err != nil {
		return err
	}

	for _, order := range orders {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			s.cacheRepo.Put(order)
		}
	}

	return nil
}
