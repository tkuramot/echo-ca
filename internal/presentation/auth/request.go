package auth

type registerUserParams struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type loginUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
