//go:build wireinject
// +build wireinject

package product

import (
	handler "course-api/internal/product/delivery/http"
	repository "course-api/internal/product/repository"
	usecase "course-api/internal/product/usecase"
	media "course-api/pkg/media/cloudinary"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.ProductHandler {
	wire.Build(
		handler.NewProductHandler,
		usecase.NewProductUsecase,
		repository.NewProductRepository,
		media.NewMediaUsecase,
	)

	return &handler.ProductHandler{}
}
