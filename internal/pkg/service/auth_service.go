package service

import (
	"context"
	"projectSwagger/ent"
	"projectSwagger/ent/user"
	"projectSwagger/internal/model"
)

type AuthService struct {
	Client *ent.Client
}

func NewAuthService(client *ent.Client) *AuthService {
	return &AuthService{Client: client}
}

func (s *AuthService) Login(payload *model.LoginUserRequest) (*ent.User, error) {
	user, err := s.Client.User.Query().
		Where(user.Username(payload.Username)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}
