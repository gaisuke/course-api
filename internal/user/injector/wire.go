//go:build wireinject
// +build wireinject

package user

import (
	handler "course-api/internal/user/delivery/http"
	repository "course-api/internal/user/repository"
	usecase "course-api/internal/user/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.UserHandler {
	wire.Build(
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		handler.NewUserHandler,
	)

	return &handler.UserHandler{}
}
