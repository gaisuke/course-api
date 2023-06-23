package middleware

import (
	"net/http"
	"os"
	"strings"

	"course-api/pkg/response"

	dto "course-api/internal/oauth/dto"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Header struct {
	Authorization string `header:"Authorization" binding:"required"`
}

func AuthJwt(ctx *gin.Context) {
	var input Header
	if err := ctx.ShouldBindHeader(&input); err != nil {
		ctx.JSON(http.StatusUnauthorized, response.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"Unauthorized",
		))
		ctx.Abort()
		return
	}

	reqToken := input.Authorization
	splitToken := strings.Split(reqToken, "Bearer ")

	if len(splitToken) != 2 {
		ctx.JSON(http.StatusUnauthorized, response.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"Unauthorized",
		))
		ctx.Abort()
		return
	}

	reqToken = splitToken[1]
	claims := &dto.MapClaimsResponse{}

	token, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"Unauthorized",
		))
		ctx.Abort()
		return
	}

	if !token.Valid {
		ctx.JSON(http.StatusUnauthorized, response.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"Unauthorized",
		))
		ctx.Abort()
		return
	}

	claims = token.Claims.(*dto.MapClaimsResponse)

	ctx.Set("user", claims)
	ctx.Next()
}
