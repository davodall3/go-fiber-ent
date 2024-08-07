package repository

import "projectSwagger/internal/model"

type UserRepository interface {
	CreateUser(body *model.UserBody) error
	GetAllUser() error
}
