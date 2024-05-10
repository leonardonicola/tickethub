package repository

import (
	"database/sql"
	"fmt"

	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/modules/user/domain"
	"github.com/leonardonicola/tickethub/internal/modules/user/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

var (
	logger *config.Logger
)

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	logger = config.NewLogger()
	return &UserRepositoryImpl{DB: db}
}

func (repo *UserRepositoryImpl) Create(user *domain.User) (*domain.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		logger.Errorf("create user - hash password: %v", err)
		return nil, fmt.Errorf("create user - hash password: %v", err)
	}

	const sqlQuery = `
	INSERT INTO users
	(id, name, surname, address, email, password)
	VALUES ($1, $2, $3, $4, $5, $6)
	`
	res, err := repo.DB.Exec(sqlQuery,
		user.ID, user.Name, user.Surname, user.Address, user.Email, hashedPass)
	if err != nil {
		logger.Errorf("create user: %v", err)
		return nil, fmt.Errorf("create user: %v", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		logger.Errorf("create user - no rows affected: %v", err)
		return nil, fmt.Errorf("create user - no rows affected: %v", err)
	}
	return user, nil
}

func (repo *UserRepositoryImpl) GetById(id string) (*dto.GetUserOutputDTO, error) {
	var user dto.GetUserOutputDTO

	const sqlQuery = `
		SELECT id, name, surname, address, email,created_at, updated_at
		FROM users
		WHERE id = $1
		AND deleted_at IS NULL
	`

	row := repo.DB.QueryRow(sqlQuery, id)
	if err := row.Scan(&user.ID, &user.Name, &user.Surname, &user.Address,
		&user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		logger.Errorf("USER(get_by_id): %v", err)
		return nil, fmt.Errorf("Error to get user by id: %v", err)
	}

	return &user, nil

}
