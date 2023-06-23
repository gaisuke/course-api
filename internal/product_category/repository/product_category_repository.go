package product_category

import (
	entity "course-api/internal/product_category/entity"
	"course-api/pkg/response"
	"course-api/pkg/utils"

	"gorm.io/gorm"
)

type ProductCategoryRepository interface {
	FindAll(offset, limit int) []entity.ProductCategory
	FindOneById(id int) (*entity.ProductCategory, *response.Error)
	Create(entity entity.ProductCategory) (*entity.ProductCategory, *response.Error)
	Update(entity entity.ProductCategory) (*entity.ProductCategory, *response.Error)
	Delete(entity entity.ProductCategory) *response.Error
}

type productCategoryRepository struct {
	db *gorm.DB
}

// Create implements ProductCategoryRepository.
func (repository *productCategoryRepository) Create(entity entity.ProductCategory) (*entity.ProductCategory, *response.Error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

// Delete implements ProductCategoryRepository.
func (repository *productCategoryRepository) Delete(entity entity.ProductCategory) *response.Error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

// FindAll implements ProductCategoryRepository.
func (repository *productCategoryRepository) FindAll(offset int, limit int) []entity.ProductCategory {
	var productCategories []entity.ProductCategory

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&productCategories)

	return productCategories
}

// FindOneById implements ProductCategoryRepository.
func (repository *productCategoryRepository) FindOneById(id int) (*entity.ProductCategory, *response.Error) {
	var productCategory entity.ProductCategory

	if err := repository.db.First(&productCategory, id).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &productCategory, nil
}

// Update implements ProductCategoryRepository.
func (repository *productCategoryRepository) Update(entity entity.ProductCategory) (*entity.ProductCategory, *response.Error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func NewProductCategoryRepository(db *gorm.DB) ProductCategoryRepository {
	return &productCategoryRepository{db}
}
