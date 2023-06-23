package admin

import (
	dto "course-api/internal/admin/dto"
	entity "course-api/internal/admin/entity"
	repository "course-api/internal/admin/repository"
	"course-api/pkg/response"

	"golang.org/x/crypto/bcrypt"
)

type AdminUsecase interface {
	FindAll(offset, limit int) []entity.Admin
	FindOneById(id int) (*entity.Admin, *response.Error)
	FindOneByEmail(email string) (*entity.Admin, *response.Error)
	Create(dto dto.AdminRequestBody) (*entity.Admin, *response.Error)
	Update(id int, dto dto.AdminRequestBody) (*entity.Admin, *response.Error)
	Delete(id int) *response.Error
	TotalCountAdmin() int64
}

type adminUsecase struct {
	repository repository.AdminRepository
}

// Create implements AdminUsecase.
func (usecase *adminUsecase) Create(dto dto.AdminRequestBody) (*entity.Admin, *response.Error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	dataAdmin := entity.Admin{
		Email:    dto.Email,
		Name:     dto.Name,
		Password: string(hashedPassword),
	}

	if dto.CreatedBy != nil {
		dataAdmin.CreatedByID = dto.CreatedBy
	}

	admin, errCreateAdmin := usecase.repository.Create(dataAdmin)
	if errCreateAdmin != nil {
		return nil, errCreateAdmin
	}

	return admin, nil
}

// Delete implements AdminUsecase.
func (usecase *adminUsecase) Delete(id int) *response.Error {
	admin, err := usecase.repository.FindOneById(id)
	if err != nil {
		return err
	}

	if err := usecase.repository.Delete(*admin); err != nil {
		return err
	}

	return nil
}

// FindAll implements AdminUsecase.
func (usecase *adminUsecase) FindAll(offset int, limit int) []entity.Admin {
	return usecase.repository.FindAll(offset, limit)
}

// FindOneByEmail implements AdminUsecase.
func (usecase *adminUsecase) FindOneByEmail(email string) (*entity.Admin, *response.Error) {
	return usecase.repository.FindOneByEmail(email)
}

// FindOneById implements AdminUsecase.
func (usecase *adminUsecase) FindOneById(id int) (*entity.Admin, *response.Error) {
	return usecase.repository.FindOneById(id)
}

// TotalCountAdmin implements AdminUsecase.
func (*adminUsecase) TotalCountAdmin() int64 {
	panic("unimplemented")
}

// Update implements AdminUsecase.
func (usecase *adminUsecase) Update(id int, dto dto.AdminRequestBody) (*entity.Admin, *response.Error) {
	admin, err := usecase.repository.FindOneById(id)
	if err != nil {
		return nil, err
	}

	admin.Name = dto.Name
	admin.Email = dto.Email
	if dto.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, &response.Error{
				Code: 500,
				Err:  err,
			}
		}
		admin.Password = string(hashedPassword)
	}

	if dto.UpdatedBy != nil {
		admin.UpdatedByID = dto.UpdatedBy
	}

	updateAdmin, err := usecase.repository.Update(*admin)
	if err != nil {
		return nil, err
	}

	return updateAdmin, nil
}

func NewAdminUsecase(repository repository.AdminRepository) AdminUsecase {
	return &adminUsecase{repository}
}
