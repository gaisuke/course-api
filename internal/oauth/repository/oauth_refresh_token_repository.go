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
func (*oauthRefreshTokenRepository) Delete(entity entity.OauthRefreshToken) *response.Error {
	panic("unimplemented")
}

// FindOneByOauthAccessTokenID implements OauthRefreshTokenRepository.
func (*oauthRefreshTokenRepository) FindOneByOauthAccessTokenID(oauthAccessTokenID int) (*entity.OauthRefreshToken, *response.Error) {
	panic("unimplemented")
}

// FindOneByToken implements OauthRefreshTokenRepository.
func (*oauthRefreshTokenRepository) FindOneByToken(token string) (*entity.OauthRefreshToken, *response.Error) {
	panic("unimplemented")
}

func NewOauthRefreshTokenRepository(db *gorm.DB) OauthRefreshTokenRepository {
	return &oauthRefreshTokenRepository{db}
}
