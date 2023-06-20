package user

type UserRequestBody struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"email"`
	Password  string `json:"password" binding:"required"`
	CreatedBy *int64 `json:"created_by"`
}