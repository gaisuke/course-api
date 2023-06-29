package discount

import (
	entity "course-api/internal/discount/entity"
	"course-api/pkg/response"
	"course-api/pkg/utils"

	"gorm.io/gorm"
)

type DiscountRepository interface {
	FindAll(offset, limit int) []entity.Discount
	FindOneById(id int) (*entity.Discount, *response.Error)
	FindOneByCode(code string) (*entity.Discount, *response.Error)
	Create(entity entity.Discount) (*entity.Discount, *response.Error)
	Update(entity entity.Discount) (*entity.Discount, *response.Error)
	Delete(entity entity.Discount) *response.Error
}

type discountRepository struct {
	db *gorm.DB
}

// Create implements DiscountRepository.
func (repository *discountRepository) Create(entity entity.Discount) (*entity.Discount, *response.Error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

// Delete implements DiscountRepository.
func (repository *discountRepository) Delete(entity entity.Discount) *response.Error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

// FindAll implements DiscountRepository.
func (repository *discountRepository) FindAll(offset int, limit int) []entity.Discount {
	var discounts []entity.Discount

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&discounts)

	return discounts
}

// FindOneByCode implements DiscountRepository.
func (repository *discountRepository) FindOneByCode(code string) (*entity.Discount, *response.Error) {
	var discount entity.Discount

	if err := repository.db.Where("code = ?", code).First(&discount).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &discount, nil
}

// FindOneById implements DiscountRepository.
func (repository *discountRepository) FindOneById(id int) (*entity.Discount, *response.Error) {
	var discount entity.Discount

	if err := repository.db.First(&discount, id).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &discount, nil
}

// Update implements DiscountRepository.
func (repository *discountRepository) Update(entity entity.Discount) (*entity.Discount, *response.Error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	return &discountRepository{db}
}
