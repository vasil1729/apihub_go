package auth

import (
	"context"
	"errors"
	"time"

	"github.com/ultimatum/apihub_go/internal/domain/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	collection *mongo.Collection
}

func NewAuthService(db *mongo.Database) *AuthService {
	return &AuthService{
		collection: db.Collection("users"),
	}
}

func (s *AuthService) Register(ctx context.Context, req auth.RegisterRequest) (*auth.User, error) {
	// Check if user exists (email or username)
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
