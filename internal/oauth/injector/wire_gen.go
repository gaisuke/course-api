// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package oauth

import (
	"course-api/internal/admin/repository"
	admin2 "course-api/internal/admin/usecase"
	"course-api/internal/oauth/delivery/http"
	oauth2 "course-api/internal/oauth/repository"
	oauth3 "course-api/internal/oauth/usecase"
	"course-api/internal/user/repository"
	"course-api/internal/user/usecase"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *oauth.OauthHandler {
	oauthClientRepository := oauth2.NewOauthClientRepository(db)
	oauthAccessTokenRepository := oauth2.NewOauthAccessTokenRepository(db)
	oauthRefreshTokenRepository := oauth2.NewOauthRefreshTokenRepository(db)
	userRepository := user.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	adminRepository := admin.NewAdminRepository(db)
	adminUsecase := admin2.NewAdminUsecase(adminRepository)
	oauthUsecase := oauth3.NewOauthUsecase(oauthClientRepository, oauthAccessTokenRepository, oauthRefreshTokenRepository, userUsecase, adminUsecase)
	oauthHandler := oauth.NewOauthHandler(oauthUsecase)
	return oauthHandler
}
