package model

import "context"

/// Service Layer: methods that handler layer expects.
type UserService interface {
	Register(ctx context.Context, user User) error
}

/// Repository layer: methods that service layer expects.
type UserRepository interface {
	Create(ctx context.Context, user User) error
}
