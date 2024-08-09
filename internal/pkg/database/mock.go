package database

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"projectSwagger/ent"
	"projectSwagger/internal/model"
)

var products = []*model.Product{
	{Name: "TV", Price: 1000.34, Quantity: 5},
	{Name: "MacBook", Price: 2500.12, Quantity: 10},
	{Name: "Lenovo", Price: 1400.45, Quantity: 4},
	{Name: "Samsung", Price: 1600.55, Quantity: 3},
	{Name: "Pixel 8", Price: 1200.67, Quantity: 6},
}

func AddProducts(client *ent.Client) {
	existedProducts, err := client.Product.
		Query().
		All(context.Background())

	needToInsert := err == nil && len(existedProducts) == 0
	if needToInsert == false {
		return
	}

	for _, product := range products {
		fmt.Println("Adding product: ", product.Name)
		_, err := client.Product.Create().
			SetName(product.Name).
			SetPrice(decimal.NewFromFloat(product.Price)).
			SetQuantity(product.Quantity).
			Save(context.Background())

		if err != nil {
			fmt.Println("Saving error", err)
			return
		}
	}

	products, err := client.Product.
		Query().
		All(context.Background())

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("result=====>>>>", products)
}
