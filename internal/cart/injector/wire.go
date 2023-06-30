//go:build wireinject
// +build wireinject

package cart

import (
	handler "course-api/internal/cart/delivery/http"
	repository "course-api/internal/cart/repository"
	usecase "course-api/internal/cart/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.CartHandler {
	wire.Build(
		repository.NewCartRepository,
		usecase.NewCartUsecase,
		handler.NewCartHandler,
	)
	return &handler.CartHandler{}
}
