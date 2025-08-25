package mapper

import (
	"fmt"
	"time"

	"github.com/paincake00/order-service/internal/delivery/dto"
	"github.com/paincake00/order-service/internal/domain/model"
)

func FromDTOToModel(dto *dto.OrderDTO) (*model.OrderModel, error) {
	createdAt, err := time.Parse(time.RFC3339, dto.DateCreated)
	if err != nil {
		return nil, fmt.Errorf("invalid date_created: %w", err)
	}

	delivery := model.DeliveryModel{
		Name:    dto.Delivery.Name,
		Phone:   dto.Delivery.Phone,
		Zip:     dto.Delivery.Zip,
		City:    dto.Delivery.City,
		Address: dto.Delivery.Address,
		Region:  dto.Delivery.Region,
	}
	if dto.Delivery.Email != nil {
		delivery.Email = *dto.Delivery.Email
	}

	payment := model.PaymentModel{
		Transaction: dto.Payment.Transaction,
		RequestID:   dto.Payment.RequestID,
		Currency:    dto.Payment.Currency,
		Provider:    dto.Payment.Provider,
		Amount:      dto.Payment.Amount,
		PaymentDT:   dto.Payment.PaymentDT,
		Bank:        dto.Payment.Bank,
		GoodsTotal:  dto.Payment.GoodsTotal,
	}
	if dto.Payment.DeliveryCost != nil {
		payment.DeliveryCost = *dto.Payment.DeliveryCost
	}
	if dto.Payment.CustomFee != nil {
		payment.CustomFee = *dto.Payment.CustomFee
	}

	items := make([]model.ItemModel, len(dto.Items))
	for i, item := range dto.Items {
		items[i] = model.ItemModel{
			TrackNumber: item.TrackNumber,
			Rid:         item.Rid,
			Name:        item.Name,
			Size:        item.Size,
			Brand:       item.Brand,
			Status:      item.Status,
		}
		if item.ChrtID != nil {
			items[i].ChrtID = *item.ChrtID
		}
		if item.Price != nil {
			items[i].Price = *item.Price
		}
		if item.Sale != nil {
			items[i].Sale = *item.Sale
		}
		if item.TotalPrice != nil {
			items[i].TotalPrice = *item.TotalPrice
		}
		if item.NmID != nil {
			items[i].NmID = *item.NmID
		}
	}

	modelVar := &model.OrderModel{
		OrderUID:          dto.OrderUID,
		TrackNumber:       dto.TrackNumber,
		Entry:             dto.Entry,
		Delivery:          delivery,
		Payment:           payment,
		Items:             items,
		Locale:            dto.Locale,
		CustomerID:        dto.CustomerID,
		DateCreated:       createdAt,
		InternalSignature: dto.InternalSignature,
		DeliveryService:   dto.DeliveryService,
	}

	if dto.ShardKey != nil {
		modelVar.ShardKey = *dto.ShardKey
	}
	if dto.SmID != nil {
		modelVar.SmID = *dto.SmID
	}
	if dto.OofShard != nil {
		modelVar.OofShard = *dto.OofShard
	}

	return modelVar, nil
}

func FromModelToDTO(model *model.OrderModel) *dto.OrderDTO {
	delivery := dto.DeliveryDTO{
		Name:    model.Delivery.Name,
		Phone:   model.Delivery.Phone,
		Zip:     model.Delivery.Zip,
		City:    model.Delivery.City,
		Address: model.Delivery.Address,
		Region:  model.Delivery.Region,
	}
	if model.Delivery.Email != "" {
		delivery.Email = &model.Delivery.Email
	}

	payment := dto.PaymentDTO{
		Transaction: model.Payment.Transaction,
		RequestID:   model.Payment.RequestID,
		Currency:    model.Payment.Currency,
		Provider:    model.Payment.Provider,
		Amount:      model.Payment.Amount,
		PaymentDT:   model.Payment.PaymentDT,
		Bank:        model.Payment.Bank,
		GoodsTotal:  model.Payment.GoodsTotal,
	}
	payment.DeliveryCost = &model.Payment.DeliveryCost
	payment.CustomFee = &model.Payment.CustomFee

	items := make([]dto.ItemDTO, len(model.Items))
	for i, item := range model.Items {
		items[i] = dto.ItemDTO{
			ChrtID:      &item.ChrtID,
			TrackNumber: item.TrackNumber,
			Price:       &item.Price,
			Rid:         item.Rid,
			Name:        item.Name,
			Sale:        &item.Sale,
			Size:        item.Size,
			TotalPrice:  &item.TotalPrice,
			NmID:        &item.NmID,
			Brand:       item.Brand,
			Status:      item.Status,
		}
	}
	
	dtoVar := &dto.OrderDTO{
		OrderUID:          model.OrderUID,
		TrackNumber:       model.TrackNumber,
		Entry:             model.Entry,
		Delivery:          delivery,
		Payment:           payment,
		Items:             items,
		Locale:            model.Locale,
		CustomerID:        model.CustomerID,
		DateCreated:       model.DateCreated.Format(time.RFC3339),
		InternalSignature: model.InternalSignature,
		DeliveryService:   model.DeliveryService,
		ShardKey:          &model.ShardKey,
		SmID:              &model.SmID,
		OofShard:          &model.OofShard,
	}

	return dtoVar
}
