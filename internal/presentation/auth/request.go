package auth

type loginUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
