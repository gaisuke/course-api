//go:build wireinject
// +build wireinject

package forgot_password

import (
	handler "course-api/internal/forgot_password/delivery/http"
	repository "course-api/internal/forgot_password/repository"
	usecase "course-api/internal/forgot_password/usecase"
	userRepository "course-api/internal/user/repository"
	userUsecase "course-api/internal/user/usecase"
	mail "course-api/pkg/mail/sendgrid"

	"github.com/google/wire"

	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.ForgotPasswordHandler {
	wire.Build(
		handler.NewForgotPasswordHandler,
		usecase.NewForgotPasswordUsecase,
		repository.NewForgotPasswordRepository,
		userUsecase.NewUserUsecase,
		userRepository.NewUserRepository,
		mail.NewMailUsecase,
	)
	return &handler.ForgotPasswordHandler{}
}
