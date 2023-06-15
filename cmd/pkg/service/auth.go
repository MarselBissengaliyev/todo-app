package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/MarselBisengaliev/go-todo-app"
	"github.com/MarselBisengaliev/go-todo-app/cmd/pkg/repository"
	"github.com/golang-jwt/jwt"
)

const (
	salt       = "hhdq8whd12hdu1dhadasjb"
	signingKey = "qweqwbdqhwbdqh1231sd123dqsaf34e5"
	tokenTTL   = 12 * time.Hour
)

type tokenClamis struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClamis{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
