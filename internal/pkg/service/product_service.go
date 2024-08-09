package service

import (
	"context"
	"html/template"
	"projectSwagger/ent"
	"projectSwagger/ent/product"
	"projectSwagger/ent/user"
	"projectSwagger/internal/model"
)

type ProductService struct {
	Client *ent.Client
}

func NewProductService(client *ent.Client) *ProductService {
	return &ProductService{Client: client}
}

func (p *ProductService) GetAllProducts() ([]*ent.Product, error) {
	products, err := p.Client.Product.
		Query().
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductService) BuyProduct(payload *model.BuyProductBody) (*ent.User, error) {
	ctx := context.Background()

	userData, userError := p.Client.User.
		Query().
		Where(user.ID(payload.UserId)).
		Only(ctx)

	if userError != nil {
		return nil, userError
	}

	productData, productError := p.Client.Product.
		Query().
		Where(product.ID(payload.ProductId)).
		Only(ctx)

	if productError != nil {
		return nil, productError
	}

	difference := userData.Balance.Sub(productData.Price)

	if difference.IsNegative() {
		return nil, &template.Error{Name: "The user's balance is insufficient"}
	}

	updatedUser, err := p.Client.User.
		UpdateOneID(userData.ID).
		SetBalance(difference).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	_, err = p.Client.Product.
		UpdateOneID(updatedUser.ID).
		SetQuantity(productData.Quantity - 1).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
