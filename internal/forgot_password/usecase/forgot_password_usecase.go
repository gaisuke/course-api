package forgot_password

import (
	dto "course-api/internal/forgot_password/dto"
	entity "course-api/internal/forgot_password/entity"
	repository "course-api/internal/forgot_password/repository"
	userDto "course-api/internal/user/dto"
	userUsecase "course-api/internal/user/usecase"
	mail "course-api/pkg/mail/sendgrid"
	response "course-api/pkg/response"
	"course-api/pkg/utils"
	"errors"
	"time"
)

type ForgotPasswordUsecase interface {
	Create(dtoForgotPassword dto.ForgotPasswordRequestBody) (*entity.ForgotPassword, *response.Error)
	Update(dto dto.ForgotPasswordUpdateRequestBody) (*entity.ForgotPassword, *response.Error)
}

type forgotPasswordUsecase struct {
	repository  repository.ForgotPasswordRepository
	userUsecase userUsecase.UserUsecase
	mail        mail.Mail
}

// Create implements ForgotPasswordUsecase.
func (usecase *forgotPasswordUsecase) Create(dtoForgotPassword dto.ForgotPasswordRequestBody) (*entity.ForgotPassword, *response.Error) {
	user, err := usecase.userUsecase.FindByEmail(dtoForgotPassword.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &response.Error{
			Code: 404,
			Err:  errors.New("email tidak ditemukan"),
		}
	}

	dateTime := time.Now().Add(time.Hour * 24)

	forgotPassword := entity.ForgotPassword{
		UserID:    &user.ID,
		Valid:     true,
		Code:      utils.RandString(32),
		ExpiredAt: &dateTime,
	}

	dataForgotPassword, err := usecase.repository.Create(forgotPassword)

	dataEmailForgotPassword := dto.ForgotPasswordEmailRequestBody{
		SUBJECT: "Forgot Password",
		EMAIL:   user.Email,
		CODE:    forgotPassword.Code,
	}

	go usecase.mail.SendForgotPassword(user.Email, dataEmailForgotPassword)
	if err != nil {
		return nil, err
	}

	return dataForgotPassword, nil

}

// Update implements ForgotPasswordUsecase.
func (usecase *forgotPasswordUsecase) Update(dto dto.ForgotPasswordUpdateRequestBody) (*entity.ForgotPassword, *response.Error) {
	code, err := usecase.repository.FindOneByCode(dto.Code)
	if err != nil || !code.Valid {
		return nil, &response.Error{
			Code: 400,
			Err:  errors.New("code tidak valid"),
		}
	}

	user, err := usecase.userUsecase.FindOneById(int(*code.UserID))
	if err != nil {
		return nil, err
	}

	dataUser := userDto.UserUpdateRequestBody{
		Password: &dto.Password,
	}

	_, err = usecase.userUsecase.Update(int(user.ID), dataUser)
	if err != nil {
		return nil, err
	}

	code.Valid = false

	usecase.repository.Update(*code)

	return code, nil
}

func NewForgotPasswordUsecase(repository repository.ForgotPasswordRepository, userUsecase userUsecase.UserUsecase, mail mail.Mail) ForgotPasswordUsecase {
	return &forgotPasswordUsecase{repository, userUsecase, mail}
}
