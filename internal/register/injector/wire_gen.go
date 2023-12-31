// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package register

import (
	"course-api/internal/register/delivery/http"
	register2 "course-api/internal/register/usecase"
	"course-api/internal/user/repository"
	user2 "course-api/internal/user/usecase"
	"course-api/pkg/mail/sendgrid"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *register.RegisterHandler {
	userRepository := user.NewUserRepository(db)
	userUsecase := user2.NewUserUsecase(userRepository)
	mailMail := mail.NewMailUsecase()
	registerUsecase := register2.NewRegisterUsecase(userUsecase, mailMail)
	registerHandler := register.NewRegisterHandler(registerUsecase)
	return registerHandler
}
