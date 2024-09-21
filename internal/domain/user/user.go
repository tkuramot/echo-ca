package user

import (
	"net/mail"
	"unicode/utf8"

	"github/tkuramot/echo-practice/pkg/ulid"

	errDomain "github/tkuramot/echo-practice/internal/domain/error"
)

type User struct {
	id       string
	email    string
	nickname string
}

const (
	nameLengthMin = 3
	nameLengthMax = 255
)

func Reconstruct(id, email, nickname string) (*User, error) {
	return newUser(id, email, nickname)
}

func NewUser(email, nickname string) (*User, error) {
	return newUser(ulid.NewULID(), email, nickname)
}

func newUser(id, email, nickname string) (*User, error) {
	if utf8.RuneCountInString(nickname) < nameLengthMin || utf8.RuneCountInString(nickname) > nameLengthMax {
		return nil, errDomain.NewError("nickname value is invalid")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errDomain.NewError("email value is invalid")
	}

	return &User{
		id:       id,
		email:    email,
		nickname: nickname,
	}, nil
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Nickname() string {
	return u.nickname
}
