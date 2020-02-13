package application

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

const AuthTokenKey = "auth-token"

func Authentication(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")

	if err != nil {
		log.Fatal(err)

		return nil, status.Errorf(
			codes.Unauthenticated,
			"could not read auth token: %v",
			err,
		)
	}

	return setAuthToken(ctx, token), nil
}

func setAuthToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, AuthTokenKey, token)
}

func GetAuthToken(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(AuthTokenKey).(string)
	return val, ok
}
