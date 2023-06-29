package user

import (
	entity "course-api/internal/user/entity"
	"course-api/pkg/response"
	"course-api/pkg/utils"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(offset int, limit int) []entity.User
	FindOneById(id int) (*entity.User, *response.Error)
	FindByEmail(email string) (*entity.User, *response.Error)
	Create(entity entity.User) (*entity.User, *response.Error)
	FindOneByCodeVerified(codeVerified string) (*entity.User, *response.Error)
	Update(entity entity.User) (*entity.User, *response.Error)
	Delete(entity entity.User) *response.Error
	TotalCountUser() int64
}

type userRepository struct {
	db *gorm.DB
}

// Create implements UserRepository.
func (repository *userRepository) Create(entity entity.User) (*entity.User, *response.Error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

// Delete implements UserRepository.
func (repository *userRepository) Delete(entity entity.User) *response.Error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Error{
			Code: 500,
			Err:  err,
		}
	}
	return nil
}

// FindAll implements UserRepository.
func (repository *userRepository) FindAll(offset int, limit int) []entity.User {
	var users []entity.User

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&users)

	return users
}

// FindByEmail implements UserRepository.
func (repository *userRepository) FindByEmail(email string) (*entity.User, *response.Error) {
	var user entity.User

	if err := repository.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &user, nil
}

// FindOneByCodeVerified implements UserRepository.
func (repository *userRepository) FindOneByCodeVerified(codeVerified string) (*entity.User, *response.Error) {
	var user entity.User

	if err := repository.db.Where("code_verified = ?", codeVerified).First(&user).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &user, nil
}

// FindOneById implements UserRepository.
func (repository *userRepository) FindOneById(id int) (*entity.User, *response.Error) {
	var user entity.User
	if err := repository.db.First(&user, id).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}
	return &user, nil
}

// TotalCountUser implements UserRepository.
func (repository *userRepository) TotalCountUser() int64 {
	panic("unimplemented")
}

// Update implements UserRepository.
func (repository *userRepository) Update(entity entity.User) (*entity.User, *response.Error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
