package tavern

import (
	"log"

	"github.com/google/uuid"
	"github.com/roberto-carlos-tg/tavern/services/order"
)

type TavernConfiguration func(os *Tavern) error

type Tavern struct {
	// OrderService to takes orders
	OrderService *order.OrderService

	// Billing Service
	BillingService interface{}
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	t := &Tavern{}

	for _, cfgs := range cfgs {
		if err := cfgs(t); err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithOrderServices(os *order.OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}
	log.Printf("\nBill the customer: %0.0f\n", price)
	return nil
}
