package repository

import (
	"context"
	"log"
	"messaging-backend/model"

	"github.com/jmoiron/sqlx"
)

type pgUserRepository struct {
	DB *sqlx.DB
}

func NewPGUserRepository(db *sqlx.DB) model.UserRepository {
	return &pgUserRepository{
		DB: db,
	}
}

func (r *pgUserRepository) Create(ctx context.Context, user model.User) error {
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *"

	if err := r.DB.GetContext(ctx, user, query, user.Username, user.Email, user.Password); err != nil {
		log.Printf("could not create a user with email: %v. Reason: %v\n", user.Email, err)
		return err
	}

	return nil
}
