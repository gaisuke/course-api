package cart

import (
	"errors"

	dto "course-api/internal/cart/dto"
	entity "course-api/internal/cart/entity"
	repository "course-api/internal/cart/repository"
	"course-api/pkg/response"
)

type CartUsecase interface {
	FindAllByUserId(userId, offset, limit int) []entity.Cart
	FindOneById(id int) (*entity.Cart, *response.Error)
	Create(dto dto.CartRequestBody) (*entity.Cart, *response.Error)
	Delete(id int, userId int) *response.Error
	DeleteByUserId(userId int) *response.Error
	Update(id int, dto dto.CartRequestUpdateBody) (*entity.Cart, *response.Error)
}

type cartUsecase struct {
	repository repository.CartRepository
}

// Create implements CartUsecase
func (usecase *cartUsecase) Create(dto dto.CartRequestBody) (*entity.Cart, *response.Error) {
	cart := &entity.Cart{
		UserID:      &dto.UserID,
		ProductID:   &dto.ProductID,
		Quantity:    1,
		IsChecked:   true,
		CreatedByID: &dto.CreatedBy,
	}

	data, err := usecase.repository.Create(*cart)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// Delete implements CartUsecase
func (usecase *cartUsecase) Delete(id int, userId int) *response.Error {
	// Cari datanya terlebih dahulu berdasarkan id
	cart, err := usecase.repository.FindOneById(id)

	if err != nil {
		return err
	}

	user := int64(userId)

	if *cart.UserID != user {
		return &response.Error{
			Code: 400,
			Err:  errors.New("cart is not owned by user"),
		}
	}

	err = usecase.repository.Delete(*cart)

	if err != nil {
		return err
	}

	return nil
}

// DeleteByUserId implements CartUsecase
func (usecase *cartUsecase) DeleteByUserId(userId int) *response.Error {
	err := usecase.repository.DeleteByUserId(userId)

	if err != nil {
		return err
	}

	return nil
}

// FindAllByUserId implements CartUsecase
func (usecase *cartUsecase) FindAllByUserId(userId, offset, limit int) []entity.Cart {
	return usecase.repository.FindAllByUserId(userId, offset, limit)
}

// FindOneById implements CartUsecase
func (usecase *cartUsecase) FindOneById(id int) (*entity.Cart, *response.Error) {
	return usecase.repository.FindOneById(id)
}

// Update implements CartUsecase
func (usecase *cartUsecase) Update(id int, dto dto.CartRequestUpdateBody) (*entity.Cart, *response.Error) {
	cart, err := usecase.repository.FindOneById(id)

	if err != nil {
		return nil, err
	}

	if *cart.UserID != *dto.UserID {
		return nil, &response.Error{
			Code: 400,
			Err:  errors.New("cart is not owned by user"),
		}
	}

	cart.IsChecked = dto.IsChecked
	cart.UpdatedByID = dto.UserID

	updateCart, err := usecase.repository.Update(*cart)

	if err != nil {
		return nil, err
	}

	return updateCart, nil
}

func NewCartUsecase(repository repository.CartRepository) CartUsecase {
	return &cartUsecase{repository}
}
