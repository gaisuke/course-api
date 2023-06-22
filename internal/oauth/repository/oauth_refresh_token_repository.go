package oauth

import (
	entity "course-api/internal/oauth/entity"
	"course-api/pkg/response"

	"gorm.io/gorm"
)

type OauthRefreshTokenRepository interface {
	Create(entity entity.OauthRefreshToken) (*entity.OauthRefreshToken, *response.Error)
	FindOneByToken(token string) (*entity.OauthRefreshToken, *response.Error)
	FindOneByOauthAccessTokenID(oauthAccessTokenID int) (*entity.OauthRefreshToken, *response.Error)
	Delete(entity entity.OauthRefreshToken) *response.Error
}

type oauthRefreshTokenRepository struct {
	db *gorm.DB
}

// Create implements OauthRefreshTokenRepository.
func (repository *oauthRefreshTokenRepository) Create(entity entity.OauthRefreshToken) (*entity.OauthRefreshToken, *response.Error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

// Delete implements OauthRefreshTokenRepository.
func (repository *oauthRefreshTokenRepository) Delete(entity entity.OauthRefreshToken) *response.Error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

// FindOneByOauthAccessTokenID implements OauthRefreshTokenRepository.
func (repository *oauthRefreshTokenRepository) FindOneByOauthAccessTokenID(oauthAccessTokenID int) (*entity.OauthRefreshToken, *response.Error) {
	var oauthRefreshToken entity.OauthRefreshToken

	if err := repository.db.Where("oauth_access_token_id = ?", oauthAccessTokenID).First(&oauthRefreshToken).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &oauthRefreshToken, nil
}

// FindOneByToken implements OauthRefreshTokenRepository.
func (repository *oauthRefreshTokenRepository) FindOneByToken(token string) (*entity.OauthRefreshToken, *response.Error) {
	var oauthRefreshToken entity.OauthRefreshToken

	if err := repository.db.Preload("OauthAccessToken").Where("token = ?", token).First(&oauthRefreshToken).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &oauthRefreshToken, nil
}

func NewOauthRefreshTokenRepository(db *gorm.DB) OauthRefreshTokenRepository {
	return &oauthRefreshTokenRepository{db}
}
