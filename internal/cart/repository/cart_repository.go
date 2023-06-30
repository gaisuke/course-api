package cart

import (
	entity "course-api/internal/cart/entity"
	"course-api/pkg/response"
	"course-api/pkg/utils"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindAllByUserId(userId, offset, limit int) []entity.Cart
	FindOneById(id int) (*entity.Cart, *response.Error)
	Create(entity entity.Cart) (*entity.Cart, *response.Error)
	Update(entity entity.Cart) (*entity.Cart, *response.Error)
	Delete(entity entity.Cart) *response.Error
	DeleteByUserId(userId int) *response.Error
}

type cartRepository struct {
	db *gorm.DB
}

// Create implements CartRepository
func (repository *cartRepository) Create(entity entity.Cart) (*entity.Cart, *response.Error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

// Delete implements CartRepository
func (repository *cartRepository) Delete(entity entity.Cart) *response.Error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

// DeleteByUserId implements CartRepository
func (repository *cartRepository) DeleteByUserId(userId int) *response.Error {
	var cart entity.Cart

	if err := repository.db.Where("user_id = ?", userId).Delete(&cart).Error; err != nil {
		return &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

// FindAllByUserId implements CartRepository
func (repository *cartRepository) FindAllByUserId(userId, offset, limit int) []entity.Cart {
	var carts []entity.Cart

	repository.db.Scopes(utils.Paginate(offset, limit)).
		Preload("User").
		Preload("Product").
		Where("user_id = ?", userId).
		Find(&carts)

	return carts
}

// FindOneById implements CartRepository
func (repository *cartRepository) FindOneById(id int) (*entity.Cart, *response.Error) {
	var cart entity.Cart

	if err := repository.db.Preload("User").Preload("Product").Find(&cart, id).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &cart, nil
}

// Update implements CartRepository
func (repository *cartRepository) Update(entity entity.Cart) (*entity.Cart, *response.Error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}
