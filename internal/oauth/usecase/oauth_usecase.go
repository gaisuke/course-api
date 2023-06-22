package oauth

import (
	adminUsecase "course-api/internal/admin/usecase"
	dto "course-api/internal/oauth/dto"
	entity "course-api/internal/oauth/entity"
	repository "course-api/internal/oauth/repository"
	userUsecase "course-api/internal/user/usecase"
	"course-api/pkg/response"
	"course-api/pkg/utils"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type OauthUsecase interface {
	Login(dtoLoginRequest dto.LoginRequestBody) (*dto.LoginResponse, *response.Error)
	Refresh(dtoRefreshToken dto.RefreshTokenRequestBody) (*dto.LoginResponse, *response.Error)
}

type oauthUsecase struct {
	oauthClientRepository       repository.OauthClientRepository
	oauthAccessTokenRepository  repository.OauthAccessTokenRepository
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository
	userUsecase                 userUsecase.UserUsecase
	adminUsecase                adminUsecase.AdminUsecase
}

// Refresh implements OauthUsecase.
func (usecase *oauthUsecase) Refresh(dtoRefreshToken dto.RefreshTokenRequestBody) (*dto.LoginResponse, *response.Error) {
	// Cek apakah refresh token ada di database
	oauthRefreshToken, err := usecase.oauthRefreshTokenRepository.FindOneByToken(dtoRefreshToken.RefreshToken)
	if err != nil {
		return nil, err
	}

	if oauthRefreshToken.ExpiredAt.Before(time.Now()) {
		return nil, &response.Error{
			Code: 400,
			Err:  errors.New("refresh token expired"),
		}
	}

	var user dto.UserResponse

	expirationTime := time.Now().Add(24 * 365 * time.Hour)
	if *oauthRefreshToken.OauthAccessToken.OauthClientID == 2 {
		dataAdmin, _ := usecase.adminUsecase.FindOneById(int(oauthRefreshToken.UserID))

		user.ID = dataAdmin.ID
		user.Name = dataAdmin.Name
		user.Email = dataAdmin.Email
	} else {
		dataUser, _ := usecase.userUsecase.FindOneById(int(oauthRefreshToken.UserID))

		user.ID = dataUser.ID
		user.Name = dataUser.Name
		user.Email = dataUser.Email
	}

	claims := &dto.ClaimsResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	if *oauthRefreshToken.OauthAccessToken.OauthClientID == 2 {
		claims.IsAdmin = true
	}

	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, errSignedString := token.SignedString(jwtKey)
	if errSignedString != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  errSignedString,
		}
	}

	// Insert data ke table oauth_access_token
	dataOauthAccessToken := entity.OauthAccessToken{
		OauthClientID: oauthRefreshToken.OauthAccessToken.OauthClientID,
		UserID:        oauthRefreshToken.UserID,
		Token:         tokenString,
		Scope:         "*",
		ExpiredAt:     &expirationTime,
	}

	saveOauthAccessToken, err := usecase.oauthAccessTokenRepository.Create(dataOauthAccessToken)
	if err != nil {
		return nil, err
	}

	expirationTimeOauthRefreshToken := time.Now().Add(24 * 366 * time.Hour)

	// Insert data ke table oauth_refresh_token
	dataOauthRefreshToken := entity.OauthRefreshToken{
		OauthAccessTokenID: &saveOauthAccessToken.ID,
		UserID:             oauthRefreshToken.UserID,
		Token:              utils.RandString(128),
		ExpiredAt:          &expirationTimeOauthRefreshToken,
	}

	saveOauthRefreshToken, err := usecase.oauthRefreshTokenRepository.Create(dataOauthRefreshToken)
	if err != nil {
		return nil, err
	}

	// Delete refresh token lama
	errDelete := usecase.oauthRefreshTokenRepository.Delete(*oauthRefreshToken)
	if errDelete != nil {
		return nil, errDelete
	}

	// Delete access token lama
	errDelete = usecase.oauthAccessTokenRepository.Delete(*oauthRefreshToken.OauthAccessToken)
	if errDelete != nil {
		return nil, errDelete
	}

	return &dto.LoginResponse{
		AccessToken:  tokenString,
		RefreshToken: saveOauthRefreshToken.Token,
		Type:         "Bearer",
		ExpiredAt:    expirationTime.Format(time.RFC3339),
		Scope:        "*",
	}, nil
}

// Login implements OauthUsecase.
func (usecase *oauthUsecase) Login(dtoLoginRequest dto.LoginRequestBody) (*dto.LoginResponse, *response.Error) {
	oauthClient, err := usecase.oauthClientRepository.FindByClientIDAndClientSecret(dtoLoginRequest.ClientID, dtoLoginRequest.ClientSecret)
	if err != nil {
		return nil, err
	}

	var user dto.UserResponse
	if oauthClient.Name == "web-admin" {
		dataAdmin, err := usecase.adminUsecase.FindOneByEmail(dtoLoginRequest.Email)
		if err != nil {
			return nil, &response.Error{
				Code: 400,
				Err:  errors.New("invalid email or password"),
			}
		}

		user.ID = dataAdmin.ID
		user.Name = dataAdmin.Name
		user.Email = dataAdmin.Email
		user.Password = dataAdmin.Password

	} else {
		dataUser, err := usecase.userUsecase.FindByEmail(dtoLoginRequest.Email)
		if err != nil {
			return nil, &response.Error{
				Code: 400,
				Err:  errors.New("invalid email or password"),
			}
		}

		user.ID = dataUser.ID
		user.Name = dataUser.Name
		user.Email = dataUser.Email
		user.Password = dataUser.Password

	}

	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	errorBcrypt := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dtoLoginRequest.Password))

	if errorBcrypt != nil {
		return nil, &response.Error{
			Code: 400,
			Err:  errors.New("invalid email or password"),
		}
	}

	expirationTime := time.Now().Add(24 * 365 * time.Hour)

	claims := &dto.ClaimsResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	if oauthClient.Name == "web-admin" {
		claims.IsAdmin = true
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)

	// Insert data ke table oauth_access_token
	dataOauthAccessToken := entity.OauthAccessToken{
		OauthClientID: &oauthClient.ID,
		UserID:        user.ID,
		Token:         tokenString,
		Scope:         "*",
		ExpiredAt:     &expirationTime,
	}

	oauthAccessToken, err := usecase.oauthAccessTokenRepository.Create(dataOauthAccessToken)
	if err != nil {
		return nil, err
	}

	expirationTimeOauthAccessToken := time.Now().Add(24 * 366 * time.Hour)

	// Insert data ke table oauth_refresh_token
	dataOauthRefreshToken := entity.OauthRefreshToken{
		OauthAccessTokenID: &oauthAccessToken.ID,
		UserID:             user.ID,
		Token:              utils.RandString(128),
		ExpiredAt:          &expirationTimeOauthAccessToken,
	}

	oauthRefreshToken, err := usecase.oauthRefreshTokenRepository.Create(dataOauthRefreshToken)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  oauthAccessToken.Token,
		RefreshToken: oauthRefreshToken.Token,
		Type:         "Bearer",
		ExpiredAt:    expirationTime.Format(time.RFC3339),
		Scope:        "*",
	}, nil
}

func NewOauthUsecase(
	oauthClientRepository repository.OauthClientRepository,
	oauthAccessTokenRepository repository.OauthAccessTokenRepository,
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository,
	userUsecase userUsecase.UserUsecase,
	adminUsecase adminUsecase.AdminUsecase,
) OauthUsecase {
	return &oauthUsecase{
		oauthClientRepository,
		oauthAccessTokenRepository,
		oauthRefreshTokenRepository,
		userUsecase,
		adminUsecase,
	}
}
