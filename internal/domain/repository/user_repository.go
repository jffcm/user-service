package repository

import "github.com/jffcm/user-service/internal/domain/entity"

type UserRepository interface {
	Save(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	ExistsByEmail(email string) (bool, error)
}