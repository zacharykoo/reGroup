package repository

import "github.com/zacharykoo/reGroup/backend/pkg/model"

type UserRepository interface {
	Get(ID int) ([]model.User, error)
	Create(model.User) (model.User, error)
	Edit(model.User) (model.User, error)
}
