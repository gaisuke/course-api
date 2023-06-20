package register

import (
	userDto "course-api/internal/user/dto"
	userUsecase "course-api/internal/user/usecase"
	"course-api/pkg/response"
)

type RegisterUsecase interface {
	Register(dto userDto.UserRequestBody) *response.Error
}

type registerUsecase struct {
	userUsecase userUsecase.UserUsecase
}

// Register implements RegisterUsecase.
func (usecase *registerUsecase) Register(dto userDto.UserRequestBody) *response.Error {
	_, err := usecase.userUsecase.Create(dto)
	if err != nil {
		return err
	}

	// TODO: Melakukan pengiriman email verifikasi ke user - dengan SendGrid

	return nil
}

func NewRegisterUsecase(userUsecase userUsecase.UserUsecase) RegisterUsecase {
	return &registerUsecase{userUsecase}
}
