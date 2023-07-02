// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cart

import (
	"course-api/internal/cart/delivery/http"
	cart2 "course-api/internal/cart/repository"
	cart3 "course-api/internal/cart/usecase"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *cart.CartHandler {
	cartRepository := cart2.NewCartRepository(db)
	cartUsecase := cart3.NewCartUsecase(cartRepository)
	cartHandler := cart.NewCartHandler(cartUsecase)
	return cartHandler
}