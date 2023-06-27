package product

import (
	dto "course-api/internal/product/dto"
	entity "course-api/internal/product/entity"
	repository "course-api/internal/product/repository"
	media "course-api/pkg/media/cloudinary"
	"course-api/pkg/response"
)

type ProductUsecase interface {
	FindAll(offset int, limit int) []entity.Product
	FindOneById(id int) (*entity.Product, *response.Error)
	Create(dto dto.ProductRequestBody) (*entity.Product, *response.Error)
	Update(id int, dto dto.ProductRequestBody) (*entity.Product, *response.Error)
	Delete(id int) *response.Error
	TotalCountProduct() int64
}

type productUsecase struct {
	repository repository.ProductRepository
	media      media.Media
}

// TotalCountProduct implements ProductUsecase.
func (usecase *productUsecase) TotalCountProduct() int64 {
	panic("unimplemented")
}

// Create implements ProductUsecase.
func (usecase *productUsecase) Create(dto dto.ProductRequestBody) (*entity.Product, *response.Error) {
	product := entity.Product{
		ProductCategoryID: &dto.ProductCategoryID,
		Title:             dto.Title,
		Description:       &dto.Description,
		IsHighlighted:     dto.IsHighlighted,
		Price:             int64(dto.Price),
		CreatedByID:       dto.CreatedBy,
	}

	if dto.Image != nil {
		image, err := usecase.media.Upload(*dto.Image)
		if err != nil {
			return nil, err
		}

		if image != nil {
			product.Image = image
		}
	}

	if dto.Video != nil {
		video, err := usecase.media.Upload(*dto.Video)
		if err != nil {
			return nil, err
		}

		if video != nil {
			product.Video = video
		}
	}

	data, err := usecase.repository.Create(product)
	if err != nil {
		return nil, err
	}

	return data, err
}

// Delete implements ProductUsecase.
func (usecase *productUsecase) Delete(id int) *response.Error {
	product, err := usecase.repository.FindOneById(id)
	if err != nil {
		return err
	}

	err = usecase.repository.Delete(*product)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements ProductUsecase.
func (usecase *productUsecase) FindAll(offset int, limit int) []entity.Product {
	return usecase.repository.FindAll(offset, limit)
}

// FindOneById implements ProductUsecase.
func (usecase *productUsecase) FindOneById(id int) (*entity.Product, *response.Error) {
	return usecase.repository.FindOneById(id)
}

// Update implements ProductUsecase.
func (usecase *productUsecase) Update(id int, dto dto.ProductRequestBody) (*entity.Product, *response.Error) {
	product, err := usecase.repository.FindOneById(id)
	if err != nil {
		return nil, err
	}

	product.ProductCategoryID = &dto.ProductCategoryID
	product.Title = dto.Title
	product.Description = &dto.Description
	product.IsHighlighted = dto.IsHighlighted
	product.Price = int64(dto.Price)
	product.UpdatedByID = dto.UpdatedBy

	if dto.Image != nil {
		image, err := usecase.media.Upload(*dto.Image)
		if err != nil {
			return nil, err
		}

		if product.Image != nil {
			_, err := usecase.media.Delete(*product.Image)
			if err != nil {
				return nil, err
			}
		}

		if image != nil {
			product.Image = image
		}
	}

	if dto.Video != nil {
		video, err := usecase.media.Upload(*dto.Video)
		if err != nil {
			return nil, err
		}

		if product.Video != nil {
			_, err := usecase.media.Delete(*product.Video)
			if err != nil {
				return nil, err
			}
		}

		if video != nil {
			product.Video = video
		}
	}

	data, err := usecase.repository.Update(*product)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewProductUsecase(repository repository.ProductRepository, media media.Media) ProductUsecase {
	return &productUsecase{repository, media}
}
