//go:build wireinject
// +build wireinject

package oauth

import (
	handler "course-api/internal/oauth/delivery/http"
	repository "course-api/internal/oauth/repository"
	usecase "course-api/internal/oauth/usecase"
	userRepository "course-api/internal/user/repository"
	userUsecase "course-api/internal/user/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.OauthHandler {
	wire.Build(
		handler.NewOauthHandler,
		usecase.NewOauthUsecase,
		repository.NewOauthClientRepository,
		repository.NewOauthAccessTokenRepository,
		repository.NewOauthRefreshTokenRepository,
		userUsecase.NewUserUsecase,
		userRepository.NewUserRepository,
	)

	return &handler.OauthHandler{}
}
