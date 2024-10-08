package service

import (
	"context"
	"github.com/shopspring/decimal"
	"projectSwagger/ent"
	"projectSwagger/internal/model"
)

type UserService struct {
	Client *ent.Client
}

func NewUserService(client *ent.Client) *UserService {
	return &UserService{Client: client}
}

func (s UserService) CreateUser(response *model.UserBody) (*ent.User, error) {
	user, err := s.Client.User.Create().
		SetName(response.Name).
		SetSurname(response.Surname).
		SetEmail(response.Email).
		SetUsername(response.Username).
		SetPassword(response.Password).
		SetBalance(decimal.NewFromFloat(response.Balance)).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserService) GetAllUsers() ([]*ent.User, error) {
	users, err := s.Client.User.
		Query().
		All(context.Background())

	if err != nil {
		return nil, err
	}
	return users, nil
}
