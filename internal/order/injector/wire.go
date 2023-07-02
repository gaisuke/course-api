//go:build wireinject
// +build wireinject

package order

import (
	cartRepository "course-api/internal/cart/repository"
	cartUsecase "course-api/internal/cart/usecase"
	discountRepository "course-api/internal/discount/repository"
	discountUsecase "course-api/internal/discount/usecase"
	handler "course-api/internal/order/delivery/http"
	repository "course-api/internal/order/repository"
	usecase "course-api/internal/order/usecase"
	orderDetailRepository "course-api/internal/order_detail/repository"
	orderDetailUsecase "course-api/internal/order_detail/usecase"
	paymentUseCase "course-api/internal/payment/usecase"
	productRepository "course-api/internal/product/repository"
	productUsecase "course-api/internal/product/usecase"
	media "course-api/pkg/media/cloudinary"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.OrderHandler {
	wire.Build(
		handler.NewOrderHandler,
		usecase.NewOrderUsecase,
		repository.NewOrderRepository,
		cartRepository.NewCartRepository,
		cartUsecase.NewCartUsecase,
		discountUsecase.NewDiscountUsecase,
		discountRepository.NewDiscountRepository,
		productRepository.NewProductRepository,
		productUsecase.NewProductUsecase,
		orderDetailRepository.NewOrderDetailRepository,
		orderDetailUsecase.NewOrderDetailUsecase,
		paymentUseCase.NewPaymentUsecase,
		media.NewMediaUsecase,
	)

	return &handler.OrderHandler{}
}
