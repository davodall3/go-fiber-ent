package repository

import "projectSwagger/internal/app/model"

type UserRepository interface {
	CreateUser(body *model.UserBody) error
	GetAllUser() error
}
