//go:build wireinject
// +build wireinject

package register

import (
	handler "course-api/internal/register/delivery/http"
	usecase "course-api/internal/register/usecase"
	userRepository "course-api/internal/user/repository"
	userUsecase "course-api/internal/user/usecase"
	mail "course-api/pkg/mail/sendgrid"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.RegisterHandler {
	wire.Build(
		handler.NewRegisterHandler,
		usecase.NewRegisterUsecase,
		userUsecase.NewUserUsecase,
		userRepository.NewUserRepository,
		mail.NewMailUsecase,
	)

	return &handler.RegisterHandler{}
}
