package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/roberto-carlos-tg/tavern/domain/product"
	"github.com/roberto-carlos-tg/tavern/services/order"
	"github.com/roberto-carlos-tg/tavern/services/tavern"
)

func main() {
	products := productInventory()

	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		// order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)

	if err != nil {
		panic(err)
	}

	tavern, err := tavern.NewTavern(
		tavern.WithOrderServices(os),
	)

	if err != nil {
		panic(err)
	}

	uid, err := os.AddCustomer("Percy")
	if err != nil {
		panic(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)
	if err != nil {
		panic(err)
	}
}

func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "Healty Beerage", 1.99)
	if err != nil {
		panic(err)
	}

	peenuts, err := product.NewProduct("Peanuts", "Snacks", 0.9)

	if err != nil {
		panic(err)
	}

	wine, err := product.NewProduct("Wine", "nasty drink", 0.99)

	if err != nil {
		panic(err)
	}

	return []product.Product{
		beer, peenuts, wine,
	}
}
