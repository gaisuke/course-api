package oauth

import (
	entity "course-api/internal/oauth/entity"
	"course-api/pkg/response"

	"gorm.io/gorm"
)

type OauthClientRepository interface {
	FindByClientIDAndClientSecret(clientID, clientSecret string) (*entity.OauthClient, *response.Error)
}

type oauthClientRepository struct {
	db *gorm.DB
}

// FindByClientIDAndClientSecret implements OauthClientRepository.
func (repository *oauthClientRepository) FindByClientIDAndClientSecret(clientID string, clientSecret string) (*entity.OauthClient, *response.Error) {
	var oauthClient entity.OauthClient
	// if err := repository.db.Where("client_id = ? AND client_secret = ?", clientID, clientSecret).First(&oauthClient).Error; err != nil {
	if err := repository.db.Where("client_id = ?", clientID).Where("client_secret = ?", clientSecret).First(&oauthClient).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &oauthClient, nil
}

func NewOauthClientRepository(db *gorm.DB) OauthClientRepository {
	return &oauthClientRepository{db}
}
