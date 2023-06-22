//go:build wireinject
// +build wireinject

package admin

import (
	handler "course-api/internal/admin/delivery/http"
	repository "course-api/internal/admin/repository"
	usecase "course-api/internal/admin/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.AdminHandler {
	wire.Build(
		handler.NewAdminHandler,
		usecase.NewAdminUsecase,
		repository.NewAdminRepository,
	)
	return &handler.AdminHandler{}
}
