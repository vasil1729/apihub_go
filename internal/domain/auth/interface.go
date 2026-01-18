package auth

import "context"

type AuthService interface {
	Register(ctx context.Context, req RegisterRequest) (*User, error)
	Login(ctx context.Context, req LoginRequest) (string, string, *User, error)
	RefreshToken(ctx context.Context, req RefreshTokenRequest) (string, error)
}
