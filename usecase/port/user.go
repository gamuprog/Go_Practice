package port

import "Go_Practice/domain/entity"

type UserRepository interface {
	FindById(userID string) (entity.User, error)
}