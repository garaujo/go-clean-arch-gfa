package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/garaujo/go-clean-arch-gfa/domain"
	"github.com/sirupsen/logrus"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

func (m *mysqlUserRepository) fetch(query string, args ...interface{}) (users []domain.User, err error) {
	rows, err := m.Conn.Query(query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result := make([]domain.User, 0)
	for rows.Next() {
		u := domain.User{}
		err = rows.Scan(
			&u.ID,
			&u.Name,
			&u.Password,
			&u.UpdatedAt,
			&u.CreatedAt,
			&u.DeletedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, u)
	}
	return result, nil
}

// Fetch is a mysqlUserRepository method that fetches all users data
// Implements the Fetch method from domain.UserRepository interface
func (m *mysqlUserRepository) Fetch(limit int64) (res []domain.User, err error) {
	query := `SELECT id, name, email, updated_at, created_at, deleted_at
					FROM user ORDER BY created_at LIMIT ?`

	res, err = m.fetch(query, limit)
	if err != nil {
		return nil, err
	}
	return
}

// Store is a mysqlUserRepository method that creates a new user
// Implements the Store method from domain.UserRepository interface
func (m *mysqlUserRepository) Store(u *domain.User) (err error) {
	query := `INSERT INTO user SET name=?, email=?, password=?, updated_at=?, created_at, deleted_at`

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return
	}

	res, err := stmt.Exec(u.Name, u.Email, u.Password, time.Now(), nil, nil)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	u.ID = lastID
	return

}

// Update is a mysqlUserRepository method that updates a user info
// Implements the Update method from domain.UserRepository interface
func (m *mysqlUserRepository) Update(u *domain.User) (err error) {
	query := `UPDATE user set name=?, email=?, updated_at=? WHERE ID =?`

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return
	}
	res, err := stmt.Exec(u.Name, u.Email, time.Now())
	if err != nil {
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected != 1 {
		err = fmt.Errorf("Something not expected: %d rows affected", rowsAffected)
		return
	}
	return
}

// Delete is a mysqlUserRepository method that deletes a user
// Implements the Delete method from domain.UserRepository interface
func (m *mysqlUserRepository) Delete(id int64) (err error) {
	query := `DELETE FROM user WHERE id = ?`

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected != 1 {
		err = fmt.Errorf("Something not expected: %d rows affected", rowsAffected)
		return
	}
	return
}

// GetByID is a mysqlUserRepository method that get a users with a given ID
// Implements the GetByID method from domain.UserRepository interface
func (m *mysqlUserRepository) GetByID(id int64) (res domain.User, err error) {
	query := `SELECT id, name, email, updated_at, created_at, deleted_at
				FROM user WHERE ID = ?`

	list, err := m.fetch(query, id)
	if err != nil {
		return domain.User{}, nil
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

// GetByName is a mysqlUserRepository method that gets the users with a given name
// Implements the GetByName method from domain.UserRepository interface
func (m *mysqlUserRepository) GetByName(name string) (res domain.User, err error) {
	query := `SELECT id, name, email, updated_at, created_at, deleted_at
					FROM user WHERE name = ?`

	list, err := m.fetch(query, name)
	if err != nil {
		return
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}
	return
}
