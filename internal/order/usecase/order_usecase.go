package order

import (
	cartUsecase "course-api/internal/cart/usecase"
	discountEntity "course-api/internal/discount/entity"
	discountUsecase "course-api/internal/discount/usecase"
	dto "course-api/internal/order/dto"
	entity "course-api/internal/order/entity"
	repository "course-api/internal/order/repository"
	orderDetailEntity "course-api/internal/order_detail/entity"
	orderDetailUsecase "course-api/internal/order_detail/usecase"
	paymentDto "course-api/internal/payment/dto"
	paymentUsecase "course-api/internal/payment/usecase"
	productEntity "course-api/internal/product/entity"
	productUsecase "course-api/internal/product/usecase"
	"course-api/pkg/response"
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type OrderUsecase interface {
	FindAllByUserId(userId, offset, limit int) []entity.Order
	FindOneById(id int) (*entity.Order, *response.Error)
	FindOneByExternalId(externalId string) (*entity.Order, *response.Error)
	Create(dto dto.OrderRequestBody) (*entity.Order, *response.Error)
	Update(id int, dto dto.OrderRequestBody) (*entity.Order, *response.Error)
	TotalCountOrder() int64
}

type orderUsecase struct {
	repository         repository.OrderRepository
	cartUsecase        cartUsecase.CartUsecase
	discountUsecase    discountUsecase.DiscountUsecase
	productUsecase     productUsecase.ProductUsecase
	orderDetailUsecase orderDetailUsecase.OrderDetailUsecase
	paymentUsecase     paymentUsecase.PaymentUsecase
}

// Create implements OrderUsecase.
func (usecase *orderUsecase) Create(dto dto.OrderRequestBody) (*entity.Order, *response.Error) {
	price := 0
	totalPrice := 0
	description := ""

	var products []productEntity.Product

	order := entity.Order{
		UserID: &dto.UserID,
		Status: "pending",
	}

	var dataDiscount *discountEntity.Discount

	carts := usecase.cartUsecase.FindAllByUserId(int(dto.UserID), 1, 9999)
	if len(carts) == 0 {
		return nil, &response.Error{
			Code: 400,
			Err:  errors.New("there is no cart for this user"),
		}
	}

	if dto.DiscountCode != nil {
		discount, err := usecase.discountUsecase.FindOneByCode(*dto.DiscountCode)
		if err != nil {
			return nil, &response.Error{
				Code: 400,
				Err:  errors.New("discount code not found"),
			}
		}

		if discount.RemainingQuantity == 0 {
			return nil, &response.Error{
				Code: 400,
				Err:  errors.New("discount quota has been used up"),
			}
		}

		if discount.StartDate != nil && discount.EndDate != nil {
			if discount.StartDate.After(time.Now()) || discount.EndDate.Before(time.Now()) {
				return nil, &response.Error{
					Code: 400,
					Err:  errors.New("code has expired"),
				}
			}
		} else if discount.StartDate != nil {
			if discount.StartDate.After(time.Now()) {
				return nil, &response.Error{
					Code: 400,
					Err:  errors.New("code has expired"),
				}
			}
		} else if discount.EndDate != nil {
			if discount.EndDate.Before(time.Now()) {
				return nil, &response.Error{
					Code: 400,
					Err:  errors.New("code has expired"),
				}
			}
		}

		dataDiscount = discount
	}

	if len(carts) > 0 {
		for _, cart := range carts {
			product, err := usecase.productUsecase.FindOneById(int(*cart.ProductID))
			if err != nil {
				return nil, err
			}

			products = append(products, *product)
		}
	} else if dto.ProductID != nil {
		product, err := usecase.productUsecase.FindOneById(int(*dto.ProductID))
		if err != nil {
			return nil, err
		}

		products = append(products, *product)
	}

	for index, product := range products {
		price += int(product.Price)
		i := strconv.Itoa(index + 1)

		description += i + ". Product : " + product.Title + "<br/>"
	}

	totalPrice = price

	if dataDiscount != nil {
		if dataDiscount.Type == "rebate" {
			totalPrice = price - int(dataDiscount.Value)
		} else if dataDiscount.Type == "percentage" {
			totalPrice = price - (price / 100 * int(dataDiscount.Value))
		}

		order.DiscountID = &dataDiscount.ID
	}

	order.Price = int64(price)
	order.TotalPrice = int64(totalPrice)
	order.CreatedByID = &dto.UserID

	externalId := uuid.New().String()
	order.ExternalID = externalId

	data, err := usecase.repository.Create(order) // insert to order table
	if err != nil {
		return nil, err
	}

	for _, product := range products { // insert product items ke order_detail table
		orderDetail := orderDetailEntity.OrderDetail{
			ProductID:   &product.ID,
			Price:       product.Price,
			CreatedByID: order.UserID,
			OrderID:     data.ID,
		}

		usecase.orderDetailUsecase.Create(orderDetail)
	}

	dataPayment := paymentDto.PaymentRequestBody{
		ExternalID:  externalId,
		Amount:      int64(totalPrice),
		PayerEmail:  dto.Email,
		Description: description,
	}

	payment, err := usecase.paymentUsecase.Create(dataPayment)
	if err != nil {
		return nil, err
	}

	data.CheckoutLink = payment.InvoiceURL

	if dto.DiscountCode != nil {
		_, err := usecase.discountUsecase.UpdateRemainingQuantity(int(dataDiscount.ID), 1, "-")
		if err != nil {
			return nil, err
		}
	}

	err = usecase.cartUsecase.DeleteByUserId(int(dto.UserID))
	if err != nil {
		return nil, err
	}

	return data, nil
}

// FindAllByUserId implements OrderUsecase.
func (usecase *orderUsecase) FindAllByUserId(userId int, offset int, limit int) []entity.Order {
	return usecase.repository.FindAllByUserId(userId, offset, limit)
}

// FindOneByExternalId implements OrderUsecase.
func (usecase *orderUsecase) FindOneByExternalId(externalId string) (*entity.Order, *response.Error) {
	return usecase.repository.FindOneByExternalId(externalId)
}

// FindOneById implements OrderUsecase.
func (usecase *orderUsecase) FindOneById(id int) (*entity.Order, *response.Error) {
	return usecase.repository.FindOneById(id)
}

// TotalCountOrder implements OrderUsecase.
func (usecase *orderUsecase) TotalCountOrder() int64 {
	panic("unimplemented")
}

// Update implements OrderUsecase.
func (usecase *orderUsecase) Update(id int, dto dto.OrderRequestBody) (*entity.Order, *response.Error) {
	order, err := usecase.repository.FindOneById(id)
	if err != nil {
		return nil, err
	}

	order.Status = dto.Status

	updateOrder, err := usecase.repository.Update(*order)
	if err != nil {
		return nil, err
	}

	return updateOrder, nil
}

func NewOrderUsecase(
	repository repository.OrderRepository,
	cartUsecase cartUsecase.CartUsecase,
	discountUsecase discountUsecase.DiscountUsecase,
	productUsecase productUsecase.ProductUsecase,
	orderDetailUsecase orderDetailUsecase.OrderDetailUsecase,
	paymentUsecase paymentUsecase.PaymentUsecase,
) OrderUsecase {
	return &orderUsecase{
		repository,
		cartUsecase,
		discountUsecase,
		productUsecase,
		orderDetailUsecase,
		paymentUsecase,
	}
}
