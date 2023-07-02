package order

import (
	"course-api/internal/middleware"
	dto "course-api/internal/order/dto"
	usecase "course-api/internal/order/usecase"
	"course-api/pkg/response"
	"course-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	usecase usecase.OrderUsecase
}

func NewOrderHandler(usecase usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecase}
}

func (handler *OrderHandler) Route(r *gin.RouterGroup) {
	orderRouter := r.Group("/api/v1")

	orderRouter.Use(middleware.AuthJwt)
	{
		orderRouter.POST("/orders", handler.Create)
		orderRouter.GET("/orders", handler.FindAllByUserId)
		orderRouter.GET("/orders/:id", handler.FindById)
	}
}

func (handler *OrderHandler) FindAllByUserId(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	user := utils.GetCurrentUser(ctx)

	data := handler.usecase.FindAllByUserId(int(user.ID), offset, limit)

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *OrderHandler) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := handler.usecase.FindOneById(id)
	if err != nil {
		ctx.JSON(int(err.Code), response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			data,
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

func (handler *OrderHandler) Create(ctx *gin.Context) {
	var input dto.OrderRequestBody
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			nil,
		))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.UserID = user.ID
	input.Email = user.Email

	data, err := handler.usecase.Create(input)
	if err != nil {
		ctx.JSON(int(err.Code), response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			data,
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
