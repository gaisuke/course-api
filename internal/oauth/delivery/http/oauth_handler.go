package oauth

import (
	dto "course-api/internal/oauth/dto"
	usecase "course-api/internal/oauth/usecase"
	"course-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OauthHandler struct {
	usecase usecase.OauthUsecase
}

func NewOauthHandler(usecase usecase.OauthUsecase) *OauthHandler {
	return &OauthHandler{usecase}
}

func (handler *OauthHandler) Route(r *gin.RouterGroup) {
	oauthRouter := r.Group("/api/v1")
	oauthRouter.POST("/oauths", handler.Login)
	oauthRouter.POST("/oauths/refresh_token", handler.Refresh)
}

func (handler *OauthHandler) Login(ctx *gin.Context) {
	var input dto.LoginRequestBody
	if err := ctx.ShouldBindJSON(&input); err != nil { // ini error jika inputan user (request) tidak sesuai
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	// Kita panggil usecase login
	data, err := handler.usecase.Login(input)
	if err != nil { // ini jika terjadi error saat login
		ctx.JSON(int(err.Code), response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *OauthHandler) Refresh(ctx *gin.Context) {
	var input dto.RefreshTokenRequestBody
	if err := ctx.ShouldBindJSON(&input); err != nil { // ini error jika inputan user (request) tidak sesuai
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	data, err := handler.usecase.Refresh(input)
	if err != nil { // ini jika terjadi error saat refresh token
		ctx.JSON(int(err.Code), response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}
