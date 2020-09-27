package domain

import "time"

// User struct represents the user entity
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// UserUsecase interface represents the user's use cases contract
type UserUsecase interface {
	Fetch(limit int64) ([]User, error)
	Store(u *User) error
	Update(u *User) error
	Delete(id int64) error
	GetByID(id int64) (User, error)
	GetByName(n string) (User, error)
}

// UserRepository interface represents the user's repository contract
type UserRepository interface {
	Fetch(limit int64) ([]User, error)
	Store(u *User) error
	Update(u *User) error
	Delete(id int64) error
	GetByID(id int64) (User, error)
	GetByName(n string) (User, error)
}
