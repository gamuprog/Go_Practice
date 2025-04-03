package interactor

import "Go_Practice/domain/entity"

type IUserUseCase interface {
	FindById(userID string) (entity.User, error)
}
