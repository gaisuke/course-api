package forgot_password

import (
	dto "course-api/internal/forgot_password/dto"
	usecase "course-api/internal/forgot_password/usecase"
	"course-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordHandler struct {
	usecase usecase.ForgotPasswordUsecase
}

func NewForgotPasswordHandler(usecase usecase.ForgotPasswordUsecase) *ForgotPasswordHandler {
	return &ForgotPasswordHandler{usecase}
}

func (handler *ForgotPasswordHandler) Route(r *gin.RouterGroup) {
	forgotPasswordRouter := r.Group("/api/v1")
	forgotPasswordRouter.POST("/forgot_passwords", handler.Create) // ini route untuk request forgot password
	forgotPasswordRouter.PUT("/forgot_passwords", handler.Update)  // ini route untuk update password baru
}

func (handler *ForgotPasswordHandler) Create(ctx *gin.Context) {
	var input dto.ForgotPasswordRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	_, err := handler.usecase.Create(input)
	if err != nil {
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
		"Silahkan cek email anda untuk reset password",
	))
}

func (handler *ForgotPasswordHandler) Update(ctx *gin.Context) {
	var input dto.ForgotPasswordUpdateRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	_, err := handler.usecase.Update(input)
	if err != nil {
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
		"Reset password success",
	))
}
