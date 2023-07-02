package order

import (
	entity "course-api/internal/order/entity"
	"course-api/pkg/response"
	"course-api/pkg/utils"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAllByUserId(userId, offset, limit int) []entity.Order
	FindOneByExternalId(externalId string) (*entity.Order, *response.Error)
	FindOneById(id int) (*entity.Order, *response.Error)
	Create(entity entity.Order) (*entity.Order, *response.Error)
	Update(entity entity.Order) (*entity.Order, *response.Error)
	TotalCountOrder() int64
}

type orderRepository struct {
	db *gorm.DB
}

// Create implements OrderRepository.
func (repository *orderRepository) Create(entity entity.Order) (*entity.Order, *response.Error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

// FindAllByUserId implements OrderRepository.
func (repository *orderRepository) FindAllByUserId(userId int, offset int, limit int) []entity.Order {
	var orders []entity.Order
	repository.db.Scopes(utils.Paginate(offset, limit)).
		Preload("OrderDetails.Product").
		Where("user_id = ?", userId).
		Find(&orders)

	return orders
}

// FindOneByExternalId implements OrderRepository.
func (repository *orderRepository) FindOneByExternalId(externalId string) (*entity.Order, *response.Error) {
	var order entity.Order
	if err := repository.db.Preload("OrderDetails.Product").
		Where("external_id = ?", externalId).
		First(&order).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &order, nil
}

// FindOneById implements OrderRepository.
func (repository *orderRepository) FindOneById(id int) (*entity.Order, *response.Error) {
	var order entity.Order
	if err := repository.db.Preload("OrderDetails.Product").First(&order, id).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}
	return &order, nil
}

// TotalCountOrder implements OrderRepository.
func (repository *orderRepository) TotalCountOrder() int64 {
	panic("unimplemented")
}

// Update implements OrderRepository.
func (repository *orderRepository) Update(entity entity.Order) (*entity.Order, *response.Error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}
