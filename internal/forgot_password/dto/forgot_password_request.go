package forgot_password

import "time"

type ForgotPasswordRequestBody struct {
	Email string `json:"email" binding:"email"`
}

type ForgotPasswordUpdateRequestBody struct {
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ForgotPasswordCreateRequestBody struct {
	UserID    uint64     `json:"user_id"`
	Valid     bool       `json:"valid"`
	Code      string     `json:"code"`
	ExpiredAt *time.Time `json:"expired_at"`
}

type ForgotPasswordEmailRequestBody struct {
	SUBJECT string
	EMAIL   string
	CODE    string
}
