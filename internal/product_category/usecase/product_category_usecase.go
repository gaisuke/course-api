package product_category

import (
	dto "course-api/internal/product_category/dto"
	entity "course-api/internal/product_category/entity"
	repository "course-api/internal/product_category/repository"
	media "course-api/pkg/media/cloudinary"
	"course-api/pkg/response"
)

type ProductCategoryUsecase interface {
	FindAll(offset int, limit int) []entity.ProductCategory
	FindOneById(id int) (*entity.ProductCategory, *response.Error)
	Create(dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, *response.Error)
	Update(id int, dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, *response.Error)
	Delete(id int) *response.Error
}

type productCategoryUsecase struct {
	repository repository.ProductCategoryRepository
	media      media.Media
}

// Create implements ProductCategoryUsecase.
func (usecase *productCategoryUsecase) Create(dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, *response.Error) {
	entity := entity.ProductCategory{
		Name:        dto.Name,
		CreatedByID: dto.CreatedBy,
	}

	if dto.Image != nil {
		image, err := usecase.media.Upload(*dto.Image)
		if err != nil { // ini error jika upload gagal
			return nil, err
		}
		if image != nil { // ini jika upload berhasil, maka image di set ke entity
			entity.Image = image
		}
	}

	data, err := usecase.repository.Create(entity)
	if err != nil { // ini error jika insert ke database gagal
		return nil, err
	}

	return data, nil
}

// Delete implements ProductCategoryUsecase.
func (usecase *productCategoryUsecase) Delete(id int) *response.Error {
	productCategory, err := usecase.repository.FindOneById(id)
	if err != nil {
		return err
	}

	if err := usecase.repository.Delete(*productCategory); err != nil {
		return err
	}

	return nil
}

// FindAll implements ProductCategoryUsecase.
func (usecase *productCategoryUsecase) FindAll(offset int, limit int) []entity.ProductCategory {
	return usecase.repository.FindAll(offset, limit)
}

// FindOneById implements ProductCategoryUsecase.
func (usecase *productCategoryUsecase) FindOneById(id int) (*entity.ProductCategory, *response.Error) {
	return usecase.repository.FindOneById(id)
}

// Update implements ProductCategoryUsecase.
func (usecase *productCategoryUsecase) Update(id int, dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, *response.Error) {
	productCategory, err := usecase.repository.FindOneById(id)
	if err != nil {
		return nil, err
	}

	productCategory.Name = dto.Name
	productCategory.UpdatedByID = dto.UpdatedBy

	if dto.Image != nil {
		image, err := usecase.media.Upload(*dto.Image)
		if err != nil { // ini error jika upload gagal
			return nil, err
		}

		if productCategory.Image != nil { // ini jika image sebelumnya tidak kosong, maka hapus image sebelumnya
			_, err := usecase.media.Delete(*productCategory.Image)
			if err != nil {
				return nil, err
			}
		}

		if image != nil { // ini jika upload berhasil, maka image di set ke entity
			productCategory.Image = image
		}
	}

	data, err := usecase.repository.Update(*productCategory)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewProductCategoryUsecase(repository repository.ProductCategoryRepository, media media.Media) ProductCategoryUsecase {
	return &productCategoryUsecase{repository, media}
}
