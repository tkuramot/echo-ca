package auth

type userResponseModel struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

type registerUserResponse struct {
	User userResponseModel `json:"user"`
}
