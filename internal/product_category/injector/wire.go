//go:build wireinject
// +build wireinject

package product_category

import (
	handler "course-api/internal/product_category/delivery/http"
	repository "course-api/internal/product_category/repository"
	usecase "course-api/internal/product_category/usecase"
	media "course-api/pkg/media/cloudinary"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.ProductCategoryHandler {
	wire.Build(
		handler.NewProductCategoryHandler,
		usecase.NewProductCategoryUsecase,
		repository.NewProductCategoryRepository,
		media.NewMediaUsecase,
	)

	return &handler.ProductCategoryHandler{}
}
