package utils

import (
	"math/rand"
	"path/filepath"

	oauthDto "course-api/internal/oauth/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RandString(length int) string {
	var letterRune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, length)

	for i := range b {
		b[i] = letterRune[rand.Intn(len(letterRune))]
	}

	return string(b)
}

func Paginate(offset, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := offset

		// jika page lebih kecil dari atau sama dengan 0, maka set page ke 1
		if page <= 0 {
			page = 1
		}

		pageSize := limit

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset = (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetCurrentUser(ctx *gin.Context) *oauthDto.MapClaimsResponse {
	user, _ := ctx.Get("user")
	return user.(*oauthDto.MapClaimsResponse)
}

func GetFileName(filename string) string {
	file := filepath.Base(filename)
	return file[:len(file)-len(filepath.Ext(file))]
}
