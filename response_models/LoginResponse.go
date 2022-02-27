package response_models

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetLoginRequest() LoginRequest {
	var loginRequest LoginRequest
	return loginRequest
}
