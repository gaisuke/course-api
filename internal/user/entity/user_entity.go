package user

import (
	adminEntity "course-api/internal/admin/entity"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int64              `json:"id"`
	Name            string             `json:"name"`
	Email           string             `json:"email"`
	Password        string             `json:"-"` // - artinya tidak akan ditampilkan
	CodeVerified    string             `json:"-"`
	EmailVerifiedAt *time.Time         `json:"email_verified_at"`
	CreatedByID     *int64             `json:"created_by" gorm:"column:created_by"`
	CreatedBy       *adminEntity.Admin `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID     *int64             `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy       *adminEntity.Admin `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt       *time.Time         `json:"created_at"`
	UpdatedAt       *time.Time         `json:"updated_at"`
	DeletedAt       gorm.DeletedAt     `json:"deleted_at"`
}
