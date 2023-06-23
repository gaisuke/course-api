package middleware

import (
	"course-api/pkg/response"
	"course-api/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthAdmin(ctx *gin.Context) {
	admin := utils.GetCurrentUser(ctx)

	if !admin.IsAdmin {
		ctx.JSON(http.StatusUnauthorized, response.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"Unauthorized",
		))
		ctx.Abort()
		return
	}

	ctx.Next()
}
