package auth

import "context"

type AuthService interface {
	Register(ctx context.Context, req RegisterRequest) (*User, error)
}
