package discount

import (
	dto "course-api/internal/discount/dto"
	entity "course-api/internal/discount/entity"
	repository "course-api/internal/discount/repository"
	"course-api/pkg/response"
)

type DiscountUsecase interface {
	FindAll(offset, limit int) []entity.Discount
	FindOneById(id int) (*entity.Discount, *response.Error)
	FindOneByCode(code string) (*entity.Discount, *response.Error)
	Create(dto dto.DiscountRequestBody) (*entity.Discount, *response.Error)
	Update(id int, dto dto.DiscountRequestBody) (*entity.Discount, *response.Error)
	Delete(id int) *response.Error
	UpdateRemainingQuantity(id, quantity int, operator string) (*entity.Discount, *response.Error)
}

type discountUsecase struct {
	repository repository.DiscountRepository
}

// Create implements DiscountUsecase.
func (usecase *discountUsecase) Create(dto dto.DiscountRequestBody) (*entity.Discount, *response.Error) {
	discount := entity.Discount{
		Name:              dto.Name,
		Code:              dto.Code,
		Quantity:          dto.Quantity,
		RemainingQuantity: dto.Quantity,
		Type:              dto.Type,
		Value:             dto.Value,
		CreatedByID:       dto.CreatedBy,
	}

	if dto.StartDate != nil {
		discount.StartDate = dto.StartDate
	}

	if dto.EndDate != nil {
		discount.EndDate = dto.EndDate
	}

	data, err := usecase.repository.Create(discount)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Delete implements DiscountUsecase.
func (usecase *discountUsecase) Delete(id int) *response.Error {
	discount, err := usecase.repository.FindOneById(id)
	if err != nil {
		return err
	}

	err = usecase.repository.Delete(*discount)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements DiscountUsecase.
func (usecase *discountUsecase) FindAll(offset int, limit int) []entity.Discount {
	return usecase.repository.FindAll(offset, limit)
}

// FindOneByCode implements DiscountUsecase.
func (usecase *discountUsecase) FindOneByCode(code string) (*entity.Discount, *response.Error) {
	return usecase.repository.FindOneByCode(code)
}

// FindOneById implements DiscountUsecase.
func (usecase *discountUsecase) FindOneById(id int) (*entity.Discount, *response.Error) {
	return usecase.repository.FindOneById(id)
}

// Update implements DiscountUsecase.
func (usecase *discountUsecase) Update(id int, dto dto.DiscountRequestBody) (*entity.Discount, *response.Error) {
	discount, err := usecase.repository.FindOneById(id)
	if err != nil {
		return nil, err
	}

	discount.Name = dto.Name
	discount.Code = dto.Code
	discount.Quantity = dto.Quantity
	discount.RemainingQuantity = dto.Quantity
	discount.Type = dto.Type
	discount.Value = dto.Value

	if dto.StartDate != nil {
		discount.StartDate = dto.StartDate
	}

	if dto.EndDate != nil {
		discount.EndDate = dto.EndDate
	}

	if dto.UpdatedBy != nil {
		discount.UpdatedByID = dto.UpdatedBy
	}

	data, err := usecase.repository.Update(*discount)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateRemainingQuantity implements DiscountUsecase.
func (usecase *discountUsecase) UpdateRemainingQuantity(id, quantity int, operator string) (*entity.Discount, *response.Error) {
	discount, err := usecase.repository.FindOneById(id)
	if err != nil {
		return nil, err
	}

	if operator == "+" {
		discount.RemainingQuantity = discount.RemainingQuantity + int64(quantity)
	} else if operator == "-" {
		discount.RemainingQuantity = discount.RemainingQuantity - int64(quantity)
	}

	updateDiscount, err := usecase.repository.Update(*discount)
	if err != nil {
		return nil, err
	}

	return updateDiscount, nil
}

func NewDiscountUsecase(repository repository.DiscountRepository) DiscountUsecase {
	return &discountUsecase{repository}
}
