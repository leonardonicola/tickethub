package repository

import (
	"database/sql"
	"fmt"

	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/modules/user/domain"
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
	res, err := repo.DB.Exec("INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6)",
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
