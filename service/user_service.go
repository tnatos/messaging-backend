package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"messaging-backend/model"

	"golang.org/x/crypto/scrypt"
)

type userService struct {
	UserRepository model.UserRepository
}

func NewUserService(repository model.UserRepository) model.UserService {
	return &userService{
		UserRepository: repository,
	}
}

func (s *userService) Signup(ctx context.Context, user model.User) error {
	hashedPassword, err := hashedPassword(user.Password)

	if err != nil {
		log.Printf("Unable to signup user for email: %v\n", user.Email)
		return err
	}

	// mutate password to salted password
	user.Password = hashedPassword

	if err := s.UserRepository.Create(ctx, user) {
		return err
	}

	return nil
}

/// returns salted password
func hashedPassword(password string) (string, error) {
	// example for making salt - https://play.golang.org/p/_Aw6WeWC42I
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// using recommended cost parameters from - https://godoc.org/golang.org/x/crypto/scrypt
	shash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}

	// return hex-encoded string with salt appended to password
	// password.salt
	hashedPassword := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))

	return hashedPassword, nil
}
