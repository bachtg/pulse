package application

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/bachtg/pulse/iam/internal/domain"
)

type LoginWithGoogleResult struct {
	AccessToken string
	UserId      string
	IsNew       bool
}

func (s *IAMService) LoginWithGoogle(
	ctx context.Context,
	idToken string,
) (*LoginWithGoogleResult, error) {

	claims, err := s.GoogleVerify.Verify(ctx, idToken)
	if err != nil {
		return nil, err
	}

	var isNew bool

	user, err := s.UserRepo.FindByGoogleId(claims.GoogleId)
	if err != nil || user == nil {
		user = &domain.User{
			Id:        uuid.NewString(),
			Email:     claims.Email,
			GoogleId:  claims.GoogleId,
			CreatedAt: time.Now().UTC().Unix(),
		}

		if err := s.UserRepo.Create(user); err != nil {
			return nil, err
		}

		isNew = true
	}

	token, err := s.TokenService.Generate(user.Id)
	if err != nil {
		return nil, err
	}

	return &LoginWithGoogleResult{
		AccessToken: token,
		UserId:      user.Id,
		IsNew:       isNew,
	}, nil
}
