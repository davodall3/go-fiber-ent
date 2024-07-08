package database

import (
	"context"
	"fmt"
	"projectSwagger/ent"
	"projectSwagger/internal/app/model"
)

var products = []*model.Product{
	{Name: "TV", Price: 1000.0},
	{Name: "MacBook", Price: 2500.0},
	{Name: "Lenovo", Price: 1400.0},
	{Name: "Samsung", Price: 1600.0},
	{Name: "Pixel 8", Price: 1200.0},
}

func AddProducts(client *ent.Client) {
	for _, product := range products {
		fmt.Println("Adding product: ", product.Name)
		_, err := client.Product.Create().
			SetName(product.Name).
			SetPrice(product.Price).
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
