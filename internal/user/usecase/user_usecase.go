package user

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
	Update(id int, dto dto.UserUpdateRequestBody) (*entity.User, *response.Error)
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

	if dto.CreatedBy != nil {
		user.CreatedByID = dto.CreatedBy
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
func (usecase *userUsecase) Delete(id int) *response.Error {
	user, err := usecase.repository.FindOneById(id)
	if err != nil {
		return err
	}

	err = usecase.repository.Delete(*user)
	if err != nil {
		return &response.Error{
			Code: 500,
			Err:  err.Err,
		}
	}

	return nil
}

// FindAll implements UserUsecase.
func (usecase *userUsecase) FindAll(offset int, limit int) []entity.User {
	return usecase.repository.FindAll(offset, limit)
}

// FindByEmail implements UserUsecase.
func (usecase *userUsecase) FindByEmail(email string) (*entity.User, *response.Error) {
	return usecase.repository.FindByEmail(email)
}

// FindOneByCodeVerified implements UserUsecase.
func (usecase *userUsecase) FindOneByCodeVerified(codeVerified string) (*entity.User, *response.Error) {
	return usecase.repository.FindOneByCodeVerified(codeVerified)
}

// FindOneById implements UserUsecase.
func (usecase *userUsecase) FindOneById(id int) (*entity.User, *response.Error) {
	return usecase.repository.FindOneById(id)
}

// TotalCountUser implements UserUsecase.
func (usecase *userUsecase) TotalCountUser() int64 {
	panic("unimplemented")
}

// Update implements UserUsecase.
func (usecase *userUsecase) Update(id int, dto dto.UserUpdateRequestBody) (*entity.User, *response.Error) {
	// find by id first, then update
	user, err := usecase.repository.FindOneById(id)
	if err != nil {
		return nil, err
	}

	if user.Email != dto.Email {
		user.Email = dto.Email
	}

	if dto.Password != nil {
		hashedPassword, errHashedPassword := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)
		if errHashedPassword != nil {
			return nil, &response.Error{
				Code: 500,
				Err:  errHashedPassword,
			}
		}
		user.Password = string(hashedPassword)
	}

	if dto.UpdatedBy != nil {
		user.UpdatedByID = dto.UpdatedBy
	}

	updateUser, err := usecase.repository.Update(*user)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

func NewUserUsecase(repository repository.UserRepository) UserUsecase {
	return &userUsecase{repository}
}
