package user

import (
	"fmt"
	"net/mail"
	"regexp"
	"unicode/utf8"

	errDomain "github/tkuramot/echo-practice/internal/domain/error"
	pwd "github/tkuramot/echo-practice/pkg/password"
	"github/tkuramot/echo-practice/pkg/ulid"
)

type User struct {
	id             string
	email          string
	nickname       string
	passwordDigest string
}

const (
	nameLengthMin     = 2
	nameLengthMax     = 255
	passwordLengthMin = 8
)

var ErrDuplicateUser = errDomain.NewError("ニックネームもしくはメールアドレスが既に登録されています。")

func Reconstruct(id, email, nickname, passwordDigest string) (*User, error) {
	return newUser(id, email, nickname, passwordDigest)
}

func NewUser(email, nickname, password string) (*User, error) {
	re := regexp.MustCompile(fmt.Sprintf(`^[\x21-\x7E]{%d,}$`, passwordLengthMin)) // ASCII文字で8文字以上	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[\W_]`).MatchString(password) // 特殊文字
	if !re.MatchString(password) || !hasLower || !hasUpper || !hasDigit || !hasSpecial {
		return nil, errDomain.NewError("パスワードは8文字以上、大文字・小文字・数字・特殊文字を含む必要があります。")
	}

	passwordDigest, err := pwd.Hash(password)
	if err != nil {
		return nil, err
	}
	return newUser(ulid.NewULID(), email, nickname, passwordDigest)
}

func newUser(id, email, nickname, passwordDigest string) (*User, error) {
	if utf8.RuneCountInString(nickname) < nameLengthMin || utf8.RuneCountInString(nickname) > nameLengthMax {
		return nil, errDomain.NewError("ニックネームは2文字以上、255文字以下で入力してください。")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errDomain.NewError("メールアドレスが無効な形式です。")
	}

	return &User{
		id:             id,
		email:          email,
		nickname:       nickname,
		passwordDigest: passwordDigest,
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

func (u *User) PasswordDigest() string {
	return u.passwordDigest
}

func (u *User) Authenticate(password string) error {
	err := pwd.Verify(password, u.passwordDigest)
	if err != nil {
		return errDomain.NewError("無効な認証情報です。")
	}
	return nil
}
