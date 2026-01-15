package api

import (
	"context"

	iam "github.com/bachtg/pulse/iam/api"
	"github.com/bachtg/pulse/iam/internal/application"
)

type IAMHandler struct {
	iam.UnimplementedIAMServiceServer
	app *application.IAMService
}

func (h *IAMHandler) LoginWithGoogle(ctx context.Context, req *iam.LoginWithGoogleRequest) (*iam.LoginWithGoogleResponse, error) {
	result, err := h.app.LoginWithGoogle(ctx, req.IdToken)
	if err != nil {
		return &iam.LoginWithGoogleResponse{
			Code:    1001,
			Message: err.Error(),
		}, nil
	}

	return &iam.LoginWithGoogleResponse{
		Code:    0,
		Message: "success",
		Data: &iam.LoginWithGoogleResponse_Data{
			AccessToken: result.AccessToken,
			UserId:      result.UserId,
			IsNewUser:   result.IsNew,
		},
	}, nil
}
