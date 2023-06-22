package oauth

import (
	entity "course-api/internal/oauth/entity"
	"course-api/pkg/response"

	"gorm.io/gorm"
)

type OauthAccessTokenRepository interface {
	Create(entity entity.OauthAccessToken) (*entity.OauthAccessToken, *response.Error)
	Delete(entity entity.OauthAccessToken) *response.Error
	FindOneByAccessToken(accessToken string) (*entity.OauthAccessToken, *response.Error)
}

type oauthAccessTokenRepository struct {
	db *gorm.DB
}

// Create implements OauthAccessTokenRepository.
func (repository *oauthAccessTokenRepository) Create(entity entity.OauthAccessToken) (*entity.OauthAccessToken, *response.Error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

// Delete implements OauthAccessTokenRepository.
func (repository *oauthAccessTokenRepository) Delete(entity entity.OauthAccessToken) *response.Error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

// FindOneByAccessToken implements OauthAccessTokenRepository.
func (repository *oauthAccessTokenRepository) FindOneByAccessToken(accessToken string) (*entity.OauthAccessToken, *response.Error) {
	var oauthAccessToken entity.OauthAccessToken

	if err := repository.db.Where("access_token = ?", accessToken).First(&oauthAccessToken).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &oauthAccessToken, nil
}

func NewOauthAccessTokenRepository(db *gorm.DB) OauthAccessTokenRepository {
	return &oauthAccessTokenRepository{db}
}
