package mysql

import (
	"database/sql"

	"github.com/garaujo/go-clean-arch-gfa/domain"
	"github.com/sirupsen/logrus"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

func (m *mysqlUserRepository) fetch(query string, args ...interface{}) (users []domain.User, err error) {
	rows, err := m.Conn.QueryContext(query, args...)
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

	result = make([]domain.User,0)
	for rows.Next() {
		u := domain.User{}
		userID := int64(0)
		err = rows.Scan(
			&u.ID,
			&u.Name,
			&u.Password,
			&u.UpdatedAt,
			&u.CreatedAt
			&u.DeletedAt
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result,u)
	}
	return result,nil
}

// Fetch is a mysqlUserRepository method that fetches all users data
// Implements the Fetch method from domain.UserRepository interface
func (m *mysqlUserRepository) Fetch(limit int64) (res []domain.User, err error) {
	query := `SELECT id, name, email, updated_at, created_at, deleted_at
					FROM user ORDER BY created_at LIMIT ?`

	res, err = m.fetch(query, num)
	if err != nil {
		return nil, err
	}
	return
}
