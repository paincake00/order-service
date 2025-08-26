package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/paincake00/order-service/internal/domain/model"
)

var ErrNotFound = errors.New("order not found")

type OrderDBRepository interface {
	GetByID(ctx context.Context, orderUID string) (*model.OrderModel, error)
	//CreateOrder(ctx context.Context, order *model.OrderModel) error
}

type OrderDBRepositoryImpl struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderDBRepository {
	return &OrderDBRepositoryImpl{db}
}

func (o *OrderDBRepositoryImpl) GetByID(ctx context.Context, orderUID string) (*model.OrderModel, error) {
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{
		ReadOnly: true,
	})
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	var order model.OrderModel

	orderQuery := `
        SELECT order_uid, track_number, entry, locale, internal_signature,
               customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard
        FROM orders
        WHERE order_uid = $1
    `
	if err = tx.QueryRowContext(ctx, orderQuery, orderUID).Scan(
		&order.OrderUID,
		&order.TrackNumber,
		&order.Entry,
		&order.Locale,
		&order.InternalSignature,
		&order.CustomerID,
		&order.DeliveryService,
		&order.ShardKey,
		&order.SmID,
		&order.DateCreated,
		&order.OofShard,
	); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, fmt.Errorf("get order: %w", err)
		}
	}

	deliveryQuery := `
        SELECT name, phone, zip, city, address, region, email
        FROM deliveries
        WHERE order_uid = $1
    `
	if err = tx.QueryRowContext(ctx, deliveryQuery, orderUID).Scan(
		&order.Delivery.Name,
		&order.Delivery.Phone,
		&order.Delivery.Zip,
		&order.Delivery.City,
		&order.Delivery.Address,
		&order.Delivery.Region,
		&order.Delivery.Email,
	); err != nil {
		return nil, fmt.Errorf("get delivery: %w", err)
	}

	paymentQuery := `
        SELECT transaction, request_id, currency, provider, amount, payment_dt,
               bank, delivery_cost, goods_total, custom_fee
        FROM payments
        WHERE order_uid = $1
    `
	if err = tx.QueryRowContext(ctx, paymentQuery, orderUID).Scan(
		&order.Payment.Transaction,
		&order.Payment.RequestID,
		&order.Payment.Currency,
		&order.Payment.Provider,
		&order.Payment.Amount,
		&order.Payment.PaymentDT,
		&order.Payment.Bank,
		&order.Payment.DeliveryCost,
		&order.Payment.GoodsTotal,
		&order.Payment.CustomFee,
	); err != nil {
		return nil, fmt.Errorf("get payment: %w", err)
	}

	itemsQuery := `
        SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status
        FROM items
        WHERE order_uid = $1
    `
	rows, err := tx.QueryContext(ctx, itemsQuery, orderUID)
	if err != nil {
		return nil, fmt.Errorf("get items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item model.ItemModel
		if err = rows.Scan(
			&item.ChrtID,
			&item.TrackNumber,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmID,
			&item.Brand,
			&item.Status,
		); err != nil {
			return nil, fmt.Errorf("scan item: %w", err)
		}
		order.Items = append(order.Items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return &order, nil
}
