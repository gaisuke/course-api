package oauth

import (
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
}

type oauthUsecase struct {
	oauthClientRepository       repository.OauthClientRepository
	oauthAccessTokenRepository  repository.OauthAccessTokenRepository
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository
	userUsecase                 userUsecase.UserUsecase
}

// Login implements OauthUsecase.
func (usecase *oauthUsecase) Login(dtoLoginRequest dto.LoginRequestBody) (*dto.LoginResponse, *response.Error) {
	oauthClient, err := usecase.oauthClientRepository.FindByClientIDAndClientSecret(dtoLoginRequest.ClientID, dtoLoginRequest.ClientSecret)
	if err != nil {
		return nil, err
	}

	var user dto.UserResponse

	// fmt.Println(dtoLoginRequest.Email, " => Masuk sini 1")
	dataUser, err := usecase.userUsecase.FindByEmail(dtoLoginRequest.Email)
	if err != nil {
		return nil, &response.Error{
			Code: 400,
			Err:  errors.New("invalid email or password"),
		}
	}

	// fmt.Println(dtoLoginRequest.Email, " => Masuk sini 2")

	user.ID = dataUser.ID
	user.Name = dataUser.Name
	user.Email = dataUser.Email
	user.Password = dataUser.Password

	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	errorBcrypt := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dtoLoginRequest.Password))

	if errorBcrypt != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  errorBcrypt,
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
) OauthUsecase {
	return &oauthUsecase{
		oauthClientRepository,
		oauthAccessTokenRepository,
		oauthRefreshTokenRepository,
		userUsecase,
	}
}
