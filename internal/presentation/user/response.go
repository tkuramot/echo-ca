package user

type userResponseModel struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

type getUserResponse struct {
	User userResponseModel `json:"user"`
}

type registerUserResponse struct {
	User userResponseModel `json:"user"`
}
