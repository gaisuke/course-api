// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package forgot_password

import (
	"course-api/internal/forgot_password/delivery/http"
	forgot_password2 "course-api/internal/forgot_password/repository"
	forgot_password3 "course-api/internal/forgot_password/usecase"
	"course-api/internal/user/repository"
	"course-api/internal/user/usecase"
	"course-api/pkg/mail/sendgrid"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *forgot_password.ForgotPasswordHandler {
	forgotPasswordRepository := forgot_password2.NewForgotPasswordRepository(db)
	userRepository := user.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	mailMail := mail.NewMailUsecase()
	forgotPasswordUsecase := forgot_password3.NewForgotPasswordUsecase(forgotPasswordRepository, userUsecase, mailMail)
	forgotPasswordHandler := forgot_password.NewForgotPasswordHandler(forgotPasswordUsecase)
	return forgotPasswordHandler
}
