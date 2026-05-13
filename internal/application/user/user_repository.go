package user

import (
	"game-hub-backend/internal/domain"
)

type UserRepository interface {
	Create(User *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}
