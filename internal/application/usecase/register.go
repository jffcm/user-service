package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jffcm/user-service/internal/domain/entity"
	"github.com/jffcm/user-service/internal/domain/repository"
	"github.com/jffcm/user-service/internal/domain/service"
)

type RegisterUseCase interface {
	Execute(input *RegisterUseCaseInput) (*RegisterUseCaseOutput, error)
}

type registerUseCase struct {
	userRepository repository.UserRepository
	passwordHasher service.PasswordHasher
}

func NewRegisterUseCase(userRepository repository.UserRepository, passwordHasher service.PasswordHasher) RegisterUseCase {
	return &registerUseCase{
		userRepository: userRepository,
		passwordHasher: passwordHasher,
	}
}

func (r *registerUseCase) Execute(input *RegisterUseCaseInput) (*RegisterUseCaseOutput, error) {
	userExists, err := r.userRepository.ExistsByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	if userExists {
		return nil, errors.New("a user with this email already exists")
	}

	hashedPassword, err := r.passwordHasher.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	user := entity.NewUser(input.Name, input.Email, hashedPassword)
	if err := r.userRepository.Save(user); err != nil {
		return nil, err
	}

	output := &RegisterUseCaseOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return output, nil
}

type RegisterUseCaseInput struct {
	Name     string
	Email    string
	Password string
}

type RegisterUseCaseOutput struct {
	ID        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
