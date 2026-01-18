package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ultimatum/apihub_go/internal/domain/auth"
	"github.com/ultimatum/apihub_go/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	collection *mongo.Collection
	config     *config.Config
}

func NewAuthService(db *mongo.Database, cfg *config.Config) *AuthService {
	return &AuthService{
		collection: db.Collection("users"),
		config:     cfg,
	}
}

func (s *AuthService) Register(ctx context.Context, req auth.RegisterRequest) (*auth.User, error) {
	// ... (Existing implementation, need to ensure imports are satisfied and signature matches if changed)
	// Wait, I am overwriting the file. I need to keep Register implementation.
	// I will use replace_file_content or write entire file with both methods.
	// Writing entire file is safer to ensure imports.
	
	// Check if user exists
	var existingUser auth.User
	err := s.collection.FindOne(ctx, bson.M{
		"$or": []bson.M{
			{"email": req.Email},
			{"username": req.Username},
		},
	}).Decode(&existingUser)
	
	if err == nil {
		return nil, errors.New("username or email already exists")
	}
	if err != mongo.ErrNoDocuments {
		return nil, err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := auth.User{
		ID:        primitive.NewObjectID(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = s.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) Login(ctx context.Context, req auth.LoginRequest) (string, string, *auth.User, error) {
	var user auth.User
	err := s.collection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", "", nil, errors.New("invalid credentials")
		}
		return "", "", nil, err
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", "", nil, errors.New("invalid credentials")
	}

	// Generate Tokens
	accessToken, err := s.generateAccessToken(&user)
	if err != nil {
		return "", "", nil, err
	}
	
	refreshToken, err := s.generateRefreshToken(&user)
	if err != nil {
		return "", "", nil, err
	}

	return accessToken, refreshToken, &user, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req auth.RefreshTokenRequest) (string, error) {
	// Parse token
	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.config.RefreshTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	// Extract User ID
	sub, err := claims.GetSubject()
	if err != nil {
		return "", errors.New("invalid token subject")
	}
	
	userID, err := primitive.ObjectIDFromHex(sub)
	if err != nil {
		return "", errors.New("invalid user id in token")
	}

	// Verify user exists (optional but recommended)
	var user auth.User
	err = s.collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Generate new access token
	accessToken, err := s.generateAccessToken(&user)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *AuthService) generateAccessToken(user *auth.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID.Hex(),
		"exp": time.Now().Add(15 * time.Minute).Unix(), // Short expiry
		"iat": time.Now().Unix(),
		"type": "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.AccessTokenSecret))
}

func (s *AuthService) generateRefreshToken(user *auth.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID.Hex(),
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(), // Long expiry
		"iat": time.Now().Unix(),
		"type": "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.RefreshTokenSecret))
}
