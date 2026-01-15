package infrastructure

import "github.com/bachtg/pulse/iam/internal/domain"

type UserRepository interface {
	FindByGoogleId(googleId string) (*domain.User, error)
	Create(user *domain.User) error
}
