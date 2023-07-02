package order_detail

import (
	entity "course-api/internal/order_detail/entity"
	"course-api/pkg/response"

	"gorm.io/gorm"
)

type OrderDetailRepository interface {
	Create(entity entity.OrderDetail) (*entity.OrderDetail, *response.Error)
}

type orderDetailRepository struct {
	db *gorm.DB
}

// Create implements OrderDetailRepository.
func (repository *orderDetailRepository) Create(entity entity.OrderDetail) (*entity.OrderDetail, *response.Error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

func NewOrderDetailRepository(db *gorm.DB) OrderDetailRepository {
	return &orderDetailRepository{db}
}
