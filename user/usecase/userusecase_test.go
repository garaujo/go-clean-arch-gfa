package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/garaujo/go-clean-arch-gfa/domain"
	"github.com/garaujo/go-clean-arch-gfa/domain/mocks"
	ucase "github.com/garaujo/go-clean-arch-gfa/user/usecase"
)

func TestFetch(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	mockUser := domain.User{
		ID:       1234567890,
		Name:     "John Smith",
		Email:    "john.smith@somewhere.com",
		Password: "secret_password",
	}

	mockListUser := make([]domain.User, 0)
	mockListUser = append(mockListUser, mockUser)

	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("Fetch", mock.Anything, mock.AnythingOfType("string")).Return(mockListUser, nil).Once()
		uc := ucase.NewUserUsecase(mockUserRepository)
		limit := int64(1)
		list, err := uc.Fetch(limit)

		assert.NoError(t, err)
		assert.Len(t, list, len(mockListUser))
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockUserRepository.On("Fetch", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("Unexpected Error")).Once()

		uc := ucase.NewUserUsecase(mockUserRepository)
		limit := int64(1)
		list, err := uc.Fetch(limit)

		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	mockUser := domain.User{
		ID:       1234567890,
		Name:     "John Smith",
		Email:    "john.smith@somewhere.com",
		Password: "secret_password",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("GetByID", mock.AnythingOfType("int64"), mock.AnythingOfType("string")).Return(mockUser, nil).Once()
		u := ucase.NewUserUsecase(mockUserRepository)

		us, err := u.GetByID(mockUser.ID)

		assert.NoError(t, err)
		assert.NotNil(t, us)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockUserRepository.On("GetByID", mock.Anything).Return(domain.User{}, errors.New("Unexpected")).Once()

		u := ucase.NewUserUsecase(mockUserRepository)

		us, err := u.GetByID(mockUser.ID)

		assert.Error(t, err)
		assert.Equal(t, domain.User{}, us)
		mockUserRepository.AssertExpectations(t)
	})

}

func TestStore(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	mockUser := domain.User{
		//ID:        1234567890,
		Name:      "John Smith",
		Email:     "john.smith@somewhere.com",
		Password:  "secret_password",
		UpdatedAt: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.UTC),
		DeletedAt: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
	}

	t.Run("success", func(t *testing.T) {
		tempMockUser := mockUser
		tempMockUser.ID = 0
		mockUserRepository.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(domain.User{}, domain.ErrNotFound).Once()
		//		mockUserRepository.On("Store", mock.AnythingOfType("*domain.User")).Return(nil).Once()
		mockUserRepository.On("Store", mock.AnythingOfType("*domain.User")).Return(nil).Once()

		uc := ucase.NewUserUsecase(mockUserRepository)

		err := uc.Store(&tempMockUser)

		assert.NoError(t, err)
		assert.Equal(t, mockUser.ID, tempMockUser.ID)
		assert.Equal(t, mockUser.Name, tempMockUser.Name)
		assert.Equal(t, mockUser.Email, tempMockUser.Email)
		assert.Equal(t, mockUser.Password, tempMockUser.Password)
		mockUserRepository.AssertExpectations(t)

	})

	t.Run("existing-name", func(t *testing.T) {
		existingUser := mockUser
		mockUserRepository.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(existingUser, nil).Once()

		uc := ucase.NewUserUsecase(mockUserRepository)

		err := uc.Store(&mockUser)

		assert.Error(t, err)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	mockUser := domain.User{
		ID:        1234567890,
		Name:      "John Smith",
		Email:     "john.smith@somewhere.com",
		Password:  "secret_password",
		UpdatedAt: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.UTC),
		DeletedAt: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("Update", &mockUser).Return(nil).Once()

		uc := ucase.NewUserUsecase(mockUserRepository)

		err := uc.Update(&mockUser)
		assert.NoError(t, err)
		assert.NotEqual(t, mockUser.UpdatedAt, time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC))
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("not-exists", func(t *testing.T) {
		mockUserRepository.On("Update", &mockUser).Once().Return(domain.ErrConflict)
		mockUser.ID = 0

		uc := ucase.NewUserUsecase(mockUserRepository)

		err := uc.Update(&mockUser)
		assert.Error(t, err)
		mockUserRepository.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	mockUser := domain.User{
		ID:        1234567890,
		Name:      "John Smith",
		Email:     "john.smith@somewhere.com",
		Password:  "secret_password",
		UpdatedAt: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2020, 12, 31, 23, 59, 59, 0, time.UTC),
		DeletedAt: time.Date(2020, 11, 01, 01, 01, 01, 0, time.UTC),
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepository.On("GetByID", mock.AnythingOfType("int64")).Return(mockUser, nil).Once()
		mockUserRepository.On("Delete", mock.AnythingOfType("int64")).Return(nil).Once()

		uc := ucase.NewUserUsecase(mockUserRepository)

		err := uc.Delete(mockUser.ID)

		assert.NoError(t, err)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("user-does-not-exists", func(t *testing.T) {
		mockUserRepository.On("GetByID", mock.AnythingOfType("int64")).Return(domain.User{}, nil).Once()
		//mockUserRepository.On("Delete", mock.AnythingOfType("int64")).Once().Return(domain.ErrNotFound)

		uc := ucase.NewUserUsecase(mockUserRepository)

		err := uc.Delete(mockUser.ID)

		assert.Error(t, err)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error-happens-in-db", func(t *testing.T) {
		mockUserRepository.On("GetByID", mock.AnythingOfType("int64")).Return(domain.User{}, errors.New("Unexpected Error")).Once()

		uc := ucase.NewUserUsecase(mockUserRepository)

		err := uc.Delete(mockUser.ID)

		assert.Error(t, err)
		mockUserRepository.AssertExpectations(t)
	})
}
