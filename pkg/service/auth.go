package service

import (
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/DwarfWizzard/todo-app/pkg/models"
	"github.com/DwarfWizzard/todo-app/pkg/repository"
)

var salt string = os.Getenv("HASH_SALT")

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password, user.Username)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := err
}

func generatePasswordHash(password string, username string) string {
	hash_username := sha1.New()
	hash_password := sha1.New()
	hash_username.Write([]byte(username))
	hash_password.Write([]byte(password))

	return fmt.Sprintf("%x", hash_password.Sum(hash_username.Sum([]byte(salt))))
}