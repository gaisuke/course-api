package cart

import (
	dto "course-api/internal/cart/dto"
	usecase "course-api/internal/cart/usecase"
	"course-api/internal/middleware"
	"course-api/pkg/response"
	"course-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	usecase usecase.CartUsecase
}

func NewCartHandler(usecase usecase.CartUsecase) *CartHandler {
	return &CartHandler{usecase}
}

func (handler *CartHandler) Route(r *gin.RouterGroup) {
	cartRouter := r.Group("/api/v1")

	cartRouter.Use(middleware.AuthJwt)
	{
		cartRouter.GET("/carts", handler.FindByUserId)
		cartRouter.POST("/carts", handler.Create)
		cartRouter.PATCH("/carts/:id", handler.Update)
		cartRouter.DELETE("/carts/:id", handler.Delete)
	}
}

func (handler *CartHandler) FindByUserId(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	user := utils.GetCurrentUser(ctx)

	data := handler.usecase.FindAllByUserId(int(user.ID), limit, offset)

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *CartHandler) Create(ctx *gin.Context) {
	var input dto.CartRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.UserID = user.ID
	input.CreatedBy = user.ID

	data, err := handler.usecase.Create(input)
	if err != nil {
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
		data,
	))
}

func (handler *CartHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user := utils.GetCurrentUser(ctx)

	err := handler.usecase.Delete(id, int(user.ID))
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
		"success delete cart",
	))
}

func (handler *CartHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.CartRequestUpdateBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.UserID = &user.ID

	data, err := handler.usecase.Update(id, input)
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
		data,
	))
}
