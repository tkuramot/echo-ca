//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package user

import "context"

type UserRepository interface {
	FindAll(ctx context.Context) ([]*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
	Save(ctx context.Context, user *User) error
}
