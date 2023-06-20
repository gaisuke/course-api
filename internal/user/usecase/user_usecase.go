package usecase

import (
	dto "course-api/internal/user/dto"
	entity "course-api/internal/user/entity"
	repository "course-api/internal/user/repository"
	"course-api/pkg/response"
	"course-api/pkg/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUsecase interface {
	FindAll(offset int, limit int) []entity.User
	FindByEmail(email string) (*entity.User, *response.Error)
	FindOneById(id int) (*entity.User, *response.Error)
	Create(dto dto.UserRequestBody) (*entity.User, *response.Error)
	FindOneByCodeVerified(codeVerified string) (*entity.User, *response.Error)
	Update(id int, dto dto.UserRequestBody) (*entity.User, *response.Error)
	Delete(id int) *response.Error
	TotalCountUser() int64
}

type userUsecase struct {
	repository repository.UserRepository
}

// Create implements UserUsecase.
func (usecase *userUsecase) Create(dto dto.UserRequestBody) (*entity.User, *response.Error) {
	checkUser, err := usecase.repository.FindByEmail(dto.Email)
	if err != nil && !errors.Is(err.Err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if checkUser != nil {
		return nil, &response.Error{
			Code: 409,
			Err:  errors.New("email sudah terdaftar"),
		}
	}

	hashedPassword, errHashedPassword := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if errHashedPassword != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  errHashedPassword,
		}
	}

	user := entity.User{
		Name:         dto.Name,
		Email:        dto.Email,
		Password:     string(hashedPassword),
		CodeVerified: utils.RandString(32),
	}

	dataUser, err := usecase.repository.Create(user)
	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err.Err,
		}
	}

	return dataUser, nil
}

// Delete implements UserUsecase.
func (*userUsecase) Delete(id int) *response.Error {
	panic("unimplemented")
}

// FindAll implements UserUsecase.
func (*userUsecase) FindAll(offset int, limit int) []entity.User {
	panic("unimplemented")
}

// FindByEmail implements UserUsecase.
func (*userUsecase) FindByEmail(email string) (*entity.User, *response.Error) {
	panic("unimplemented")
}

// FindOneByCodeVerified implements UserUsecase.
func (*userUsecase) FindOneByCodeVerified(codeVerified string) (*entity.User, *response.Error) {
	panic("unimplemented")
}

// FindOneById implements UserUsecase.
func (*userUsecase) FindOneById(id int) (*entity.User, *response.Error) {
	panic("unimplemented")
}

// TotalCountUser implements UserUsecase.
func (*userUsecase) TotalCountUser() int64 {
	panic("unimplemented")
}

// Update implements UserUsecase.
func (*userUsecase) Update(id int, dto dto.UserRequestBody) (*entity.User, *response.Error) {
	panic("unimplemented")
}

func NewUserUsecase(repository repository.UserRepository) UserUsecase {
	return &userUsecase{repository}
}