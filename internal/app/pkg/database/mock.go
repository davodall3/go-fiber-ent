package database

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"projectSwagger/ent"
	"projectSwagger/internal/app/model"
)

var products = []*model.Product{
	{Name: "TV", Price: 1000.34},
	{Name: "MacBook", Price: 2500.12},
	{Name: "Lenovo", Price: 1400.45},
	{Name: "Samsung", Price: 1600.55},
	{Name: "Pixel 8", Price: 1200.67},
}

func AddProducts(client *ent.Client) {
	for _, product := range products {
		fmt.Println("Adding product: ", product.Name)
		_, err := client.Product.Create().
			SetName(product.Name).
			SetPrice(decimal.NewFromFloat(product.Price)).
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
