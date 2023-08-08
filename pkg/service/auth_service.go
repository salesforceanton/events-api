package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/salesforceanton/events-api/config"
	"github.com/salesforceanton/events-api/domain"
	"github.com/salesforceanton/events-api/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
	cfg  *config.Config
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization, cfg *config.Config) *AuthService {
	return &AuthService{
		repo: repo,
		cfg:  cfg,
	}
}

func (s *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	passwordSecret := s.cfg.PasswordHashSalt

	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(passwordSecret)))
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	userId, err := s.GetUserId(username, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})

	return token.SignedString([]byte(s.cfg.TokenSecret))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&TokenClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Invalid Signing Method")
			}
			return []byte(s.cfg.TokenSecret), nil
		},
	)

	if err != nil {
		return 0, errors.New("Error with parsing Access Token")
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("Error with parsing Access Token")
	}

	return claims.UserId, nil
}

func (s *AuthService) GetUserId(username, password string) (int, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))

	if err != nil {
		return 0, errors.New("No registered User with this credentials")
	}

	return user.Id, nil
}
