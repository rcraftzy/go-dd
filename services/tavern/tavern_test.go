package tavern

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/roberto-carlos-tg/tavern/services/order"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderServices(os))
	if err != nil {
		t.Fatal(err)
	}

	uid, err := os.AddCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)

	if err != nil {
		t.Fatal(err)
	}
}
