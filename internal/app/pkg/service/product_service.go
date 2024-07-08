package service

import (
	"context"
	"html/template"
	"projectSwagger/ent"
	"projectSwagger/ent/product"
	"projectSwagger/ent/user"
	"projectSwagger/internal/app/model"
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
	u, userError := p.Client.User.
		Query().
		Where(user.ID(payload.UserId)).
		Only(context.Background())

	if userError != nil {
		return nil, userError
	}

	prod, productError := p.Client.Product.
		Query().
		Where(product.ID(payload.ProductId)).
		Only(context.Background())

	if productError != nil {
		return nil, productError
	}

	hasBalance := u.Balance - prod.Price

	if hasBalance < 0 {
		return nil, &template.Error{Name: "The user's balance is insufficient"}
	}

	updatedUser, err := p.Client.User.
		UpdateOneID(u.ID).
		SetBalance(hasBalance).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
