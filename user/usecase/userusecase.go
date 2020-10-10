package usecase

import (
	"time"

	"github.com/garaujo/go-clean-arch-gfa/domain"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

// NewUserUsecase creates a new userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: u,
	}
}

// Fetch is a userUsecase method that fetches all users data
// Implements the Fetch method from domain.UserUsecase interface
func (uc *userUsecase) Fetch(limit int64) (res []domain.User, err error) {
	if limit == 0 {
		limit = 10
	}

	res, err = uc.userRepository.Fetch(limit)
	if err != nil {
		return nil, err
	}

	return
}

// Store is a userUsecase method that creates a new user
// Implements the Store method from domain.UserUsecase interface
func (uc *userUsecase) Store(u *domain.User) (err error) {
	user, _ := uc.GetByName(u.Name)
	if user != (domain.User{}) {
		return domain.ErrConflict
	}

	err = uc.userRepository.Store(u)
	return
}

// Update is a userUsecase method that updates user info
// Implements the Update method from domain.UserUsecase interface
func (uc *userUsecase) Update(u *domain.User) (err error) {
	u.UpdatedAt = time.Now()

	return uc.userRepository.Update(u)
}

// Delete is a userUsecase method that deletes a user
// Implements the Delete method from domain.UserUsecase interface
func (uc *userUsecase) Delete(id int64) (err error) {
	user, err := uc.userRepository.GetByID(id)
	if err != nil {
		return
	}

	if user == (domain.User{}) {
		return domain.ErrNotFound
	}

	return uc.userRepository.Delete(id)
}

// GetByID is a userUsecase method that get a users with a given ID
// Implements the GetByID method from domain.UserUsecase interface
func (uc *userUsecase) GetByID(id int64) (res domain.User, err error) {
	res, err = uc.userRepository.GetByID(id)
	if err != nil {
		return
	}
	return
}

// GetByName is a userUsecase method that get a users with a given name
// Implements the GetByName method from domain.UserUsecase interface
func (uc *userUsecase) GetByName(name string) (res domain.User, err error) {
	res, err = uc.userRepository.GetByName(name)
	if err != nil {
		return
	}
	return
}
