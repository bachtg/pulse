package application

import (
	"github.com/bachtg/pulse/iam/internal/infrastructure"
)

type IAMService struct {
	UserRepo     infrastructure.UserRepository
	GoogleVerify infrastructure.GoogleTokenVerifier
	TokenService TokenService
}

type TokenService interface {
	Generate(userId string) (string, error)
}
