package infrastructure

import "context"

type GoogleClaims struct {
	GoogleId string
	Email    string
}

type GoogleTokenVerifier interface {
	Verify(ctx context.Context, idToken string) (*GoogleClaims, error)
}
