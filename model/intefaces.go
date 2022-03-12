package model

type UserRepository struct {
/// Repository layer: methods that service layer expects.
type UserRepository interface {
	Create(ctx context.Context, user User) error
}
