package repository

import (
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
