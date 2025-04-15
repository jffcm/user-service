package usecase

import (
	"errors"

	"github.com/jffcm/user-service/internal/domain/repository"
	"github.com/jffcm/user-service/internal/domain/service"
)

type LoginUseCase interface {
	Execute(input *LoginUseCaseInput) (*LoginUseCaseOutput, error)
}

type loginUseCase struct {
	userRepository repository.UserRepository
	passwordHasher service.PasswordHasher
	tokenGenerator service.TokenGenerator
}

func NewLoginUseCase(
	userRepository repository.UserRepository,
	passwordHasher service.PasswordHasher,
	tokenGenerator service.TokenGenerator,
) LoginUseCase {
	return &loginUseCase{
		userRepository: userRepository,
		passwordHasher: passwordHasher,
		tokenGenerator: tokenGenerator,
	}
}

func (l *loginUseCase) Execute(input *LoginUseCaseInput) (*LoginUseCaseOutput, error) {
	user, err := l.userRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user with this email was not found")
	}

	isPasswordValid := l.passwordHasher.Compare(user.Password, input.Password)
	if !isPasswordValid {
		return nil, errors.New("the provided password is incorrect")
	}

	token, err := l.tokenGenerator.Generate(user.Email)
	if err != nil {
		return nil, err
	}

	return &LoginUseCaseOutput{Token: token}, nil
}

type LoginUseCaseInput struct {
	Email    string
	Password string
}

type LoginUseCaseOutput struct {
	Token string
}
