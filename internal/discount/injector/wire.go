//go:build wireinject
// +build wireinject

package discount

import (
	handler "course-api/internal/discount/delivery/http"
	repository "course-api/internal/discount/repository"
	usecase "course-api/internal/discount/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.DiscountHandler {
	wire.Build(
		repository.NewDiscountRepository,
		usecase.NewDiscountUsecase,
		handler.NewDiscountHandler,
	)

	return &handler.DiscountHandler{}
}
