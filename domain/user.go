package domain

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type UserRegisterResponse struct {
	Status string `json:"status"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type UserSessionResponse struct {
	UserId float64 `json:"user_id"`
	Email  string  `json:"email"`
	Role   string  `json:"role"`
}
