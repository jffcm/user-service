package repository

import (
	"database/sql"
	"errors"

	"github.com/jffcm/user-service/internal/domain/entity"
	"github.com/jffcm/user-service/internal/domain/repository"
)

type PostgresUserRepository struct {
	DB *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) repository.UserRepository {
	return &PostgresUserRepository{DB: db}
}

func (u *PostgresUserRepository) Save(user *entity.User) error {
	query := `
		INSERT INTO users (id, name, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := u.DB.Exec(query, user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	return err
}

func (u *PostgresUserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	query := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = $1`
	err := u.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *PostgresUserRepository) ExistsByEmail(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	err := u.DB.QueryRow(query, email).Scan(&exists)

	return exists, err
}
