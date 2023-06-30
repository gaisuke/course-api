package cart

import (
	productEntity "course-api/internal/product/entity"
	userEntity "course-api/internal/user/entity"
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID          int64                  `json:"id"`
	User        *userEntity.User       `json:"user" gorm:"foreignKey:UserID;references:ID"`
	UserID      *int64                 `json:"user_id"`
	Product     *productEntity.Product `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	ProductID   *int64                 `json:"product_id"`
	Quantity    int64                  `json:"quantity"`
	IsChecked   bool                   `json:"is_checked"`
	CreatedBy   *userEntity.User       `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedByID *int64                 `json:"created_by" gorm:"column:created_by"`
	UpdatedBy   *userEntity.User       `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	UpdatedByID *int64                 `json:"updated_by" gorm:"column:updated_by"`
	CreatedAt   *time.Time             `json:"created_at"`
	UpdatedAt   *time.Time             `json:"updated_at"`
	DeletedAt   gorm.DeletedAt         `json:"deleted_at"`
}
