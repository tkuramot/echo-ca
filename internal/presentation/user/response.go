package user

type getUserResponse struct {
	User userResponseModel `json:"user"`
}

type userResponseModel struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}
