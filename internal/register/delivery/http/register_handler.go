package register

import (
	registerUsecase "course-api/internal/register/usecase"
	userDto "course-api/internal/user/dto"
	"course-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	registerUsecase registerUsecase.RegisterUsecase
}

func NewRegisterHandler(registerUsecase registerUsecase.RegisterUsecase) *RegisterHandler {
	return &RegisterHandler{registerUsecase}
}

func (handler *RegisterHandler) Route(r *gin.RouterGroup) {
	r.POST("api/v1/registers", handler.Register)
}

func (handler *RegisterHandler) Register(ctx *gin.Context) {
	// Validasi input dari user
	var registerRequest userDto.UserRequestBody
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil { // ini error jika inputan user (request) tidak sesuai
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	err := handler.registerUsecase.Register(registerRequest)
	if err != nil { // ini jika terjadi error saat register
		ctx.JSON(int(err.Code), response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, response.Response(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		"Register success, silahkan cek email anda untuk verifikasi akun",
	))
}
