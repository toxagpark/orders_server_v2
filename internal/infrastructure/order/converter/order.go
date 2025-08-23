package converter

import (
	"WB_LVL_0_NEW/internal/domain/model"
	"WB_LVL_0_NEW/internal/infrastructure/order/dto"
)

func ToDomain(d dto.Order) model.Order {
	return model.Order{
		OrderUID:    d.OrderUID,
		TrackNumber: d.TrackNumber,
		Entry:       d.Entry,
		Delivery: model.Delivery{
			Name:    d.Delivery.Name,
			Phone:   d.Delivery.Phone,
			Zip:     d.Delivery.Zip,
			City:    d.Delivery.City,
			Address: d.Delivery.Address,
			Region:  d.Delivery.Region,
			Email:   d.Delivery.Email,
		},
		Payment: model.Payment{
			Transaction:  d.Payment.Transaction,
			RequestID:    d.Payment.RequestID,
			Currency:     d.Payment.Currency,
			Provider:     d.Payment.Provider,
			Amount:       d.Payment.Amount,
			PaymentDT:    d.Payment.PaymentDT,
			Bank:         d.Payment.Bank,
			DeliveryCost: d.Payment.DeliveryCost,
			GoodsTotal:   d.Payment.GoodsTotal,
			CustomFee:    d.Payment.CustomFee,
		},
		Items:             toDomainItems(d.Items),
		Locale:            d.Locale,
		InternalSignature: d.InternalSignature,
		CustomerID:        d.CustomerID,
		DeliveryService:   d.DeliveryService,
		ShardKey:          d.ShardKey,
		SMID:              d.SMID,
		DateCreated:       d.DateCreated,
		OOFShard:          d.OOFShard,
	}
}

func ToDTO(m model.Order) dto.Order {
	return dto.Order{
		OrderUID:    m.OrderUID,
		TrackNumber: m.TrackNumber,
		Entry:       m.Entry,
		Delivery: dto.Delivery{
			Name:    m.Delivery.Name,
			Phone:   m.Delivery.Phone,
			Zip:     m.Delivery.Zip,
			City:    m.Delivery.City,
			Address: m.Delivery.Address,
			Region:  m.Delivery.Region,
			Email:   m.Delivery.Email,
		},
		Payment: dto.Payment{
			Transaction:  m.Payment.Transaction,
			RequestID:    m.Payment.RequestID,
			Currency:     m.Payment.Currency,
			Provider:     m.Payment.Provider,
			Amount:       m.Payment.Amount,
			PaymentDT:    m.Payment.PaymentDT,
			Bank:         m.Payment.Bank,
			DeliveryCost: m.Payment.DeliveryCost,
			GoodsTotal:   m.Payment.GoodsTotal,
			CustomFee:    m.Payment.CustomFee,
		},
		Items:             toDTOItems(m.Items, m.OrderUID),
		Locale:            m.Locale,
		InternalSignature: m.InternalSignature,
		CustomerID:        m.CustomerID,
		DeliveryService:   m.DeliveryService,
		ShardKey:          m.ShardKey,
		SMID:              m.SMID,
		DateCreated:       m.DateCreated,
		OOFShard:          m.OOFShard,
	}
}

func toDomainItems(items []dto.Item) []model.Item {
	result := make([]model.Item, len(items))
	for i, item := range items {
		result[i] = model.Item{
			ChrtID:      item.ChrtID,
			TrackNumber: item.TrackNumber,
			Price:       item.Price,
			RID:         item.RID,
			Name:        item.Name,
			Sale:        item.Sale,
			Size:        item.Size,
			TotalPrice:  item.TotalPrice,
			NmID:        item.NmID,
			Brand:       item.Brand,
			Status:      item.Status,
		}
	}
	return result
}

func toDTOItems(items []model.Item, orderUID string) []dto.Item {
	result := make([]dto.Item, len(items))
	for i, item := range items {
		result[i] = dto.Item{
			ChrtID:      item.ChrtID,
			TrackNumber: item.TrackNumber,
			Price:       item.Price,
			RID:         item.RID,
			Name:        item.Name,
			Sale:        item.Sale,
			Size:        item.Size,
			TotalPrice:  item.TotalPrice,
			NmID:        item.NmID,
			Brand:       item.Brand,
			Status:      item.Status,
			OrderUID:    orderUID,
		}
	}
	return result
}
