package order

import (
	"testing"

	"github.com/google/uuid"
	"github.com/roberto-carlos-tg/tavern/domain/customer"
	"github.com/roberto-carlos-tg/tavern/domain/product"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healty Beerage", 1.99)
	if err != nil {
		t.Fatal(err)
	}

	peenuts, err := product.NewProduct("Peanuts", "Snacks", 0.9)

	if err != nil {
		t.Fatal(err)
	}

	wine, err := product.NewProduct("Wine", "nasty drink", 0.99)

	if err != nil {
		t.Fatal(err)
	}

	return []product.Product{
		beer, peenuts, wine,
	}
}

func TestOrder_NewOrderSeriv(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	cust, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)

	if err != nil {
		t.Error(err)
	}
}
