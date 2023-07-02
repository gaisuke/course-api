package order

import (
	discountEntity "course-api/internal/discount/entity"
	orderDetailEntity "course-api/internal/order_detail/entity"
	userEntity "course-api/internal/user/entity"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           int64                           `json:"id"`
	User         *userEntity.User                `json:"user" gorm:"foreignKey:UserID;references:ID"`
	UserID       *int64                          `json:"user_id"`
	OrderDetails []orderDetailEntity.OrderDetail `json:"order_details"`
	Discount     *discountEntity.Discount        `json:"discount" gorm:"foreignKey:DiscountID;references:ID"`
	DiscountID   *int64                          `json:"discount_id"`
	CheckoutLink string                          `json:"checkout_link"`
	ExternalID   string                          `json:"external_id"`
	Price        int64                           `json:"price"`
	TotalPrice   int64                           `json:"total_price"`
	Status       string                          `json:"status"`
	CreatedByID  *int64                          `json:"created_by" gorm:"column:created_by"`
	CreatedBy    *userEntity.User                `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID  *int64                          `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy    *userEntity.User                `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt    *time.Time                      `json:"created_at"`
	UpdatedAt    *time.Time                      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt                  `json:"deleted_at"`
}
