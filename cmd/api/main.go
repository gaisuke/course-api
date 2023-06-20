package main

import (
	userRepository "course-api/internal/user/repository"
	userUsecase "course-api/internal/user/usecase"
	"course-api/pkg/db/mysql"

	registerHandler "course-api/internal/register/delivery/http"
	registerUsecase "course-api/internal/register/usecase"

	oauthHandler "course-api/internal/oauth/delivery/http"
	oauthRepository "course-api/internal/oauth/repository"
	oauthUsecase "course-api/internal/oauth/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := mysql.DB()

	userRepository := userRepository.NewUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(userRepository)

	registerUsecase := registerUsecase.NewRegisterUsecase(userUsecase)
	registerHandler.NewRegisterHandler(registerUsecase).Route(&r.RouterGroup)

	oauthClientRepository := oauthRepository.NewOauthClientRepository(db)
	oauthAccessTokenRepository := oauthRepository.NewOauthAccessTokenRepository(db)
	oauthRefreshTokenRepository := oauthRepository.NewOauthRefreshTokenRepository(db)

	oauthUsecase := oauthUsecase.NewOauthUsecase(oauthClientRepository, oauthAccessTokenRepository, oauthRefreshTokenRepository, userUsecase)

	oauthHandler.NewOauthHandler(oauthUsecase).Route(&r.RouterGroup)

	r.Run()
}
