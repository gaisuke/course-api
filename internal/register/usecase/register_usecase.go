package register

import (
	registerDto "course-api/internal/register/dto"
	userDto "course-api/internal/user/dto"
	userUsecase "course-api/internal/user/usecase"
	mail "course-api/pkg/mail/sendgrid"
	"course-api/pkg/response"
)

type RegisterUsecase interface {
	Register(dto userDto.UserRequestBody) *response.Error
}

type registerUsecase struct {
	userUsecase userUsecase.UserUsecase
	mail        mail.Mail
}

// Register implements RegisterUsecase.
func (usecase *registerUsecase) Register(dto userDto.UserRequestBody) *response.Error {
	user, err := usecase.userUsecase.Create(dto)
	if err != nil {
		return err
	}

	// Melakukan pengiriman email verifikasi ke user - dengan SendGrid
	data := registerDto.EmailVerification{
		SUBJECT:           "Email Verification",
		EMAIL:             dto.Email,
		VERIFICATION_CODE: user.CodeVerified,
	}

	go usecase.mail.SendVerification(dto.Email, data)

	return nil
}

func NewRegisterUsecase(userUsecase userUsecase.UserUsecase, mail mail.Mail) RegisterUsecase {
	return &registerUsecase{userUsecase, mail}
}
