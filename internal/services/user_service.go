package services

import (
	"errors"

	"github.com/anthoz69/go-temfy/internal/domain/entities"
	"github.com/anthoz69/go-temfy/internal/domain/interfaces"

	"gorm.io/gorm"
)

type UserService struct {
	userRepo interfaces.UserRepository
}

func NewUserService(userRepo interfaces.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(user *entities.User) error {
	existingUser, err := s.userRepo.GetByEmail(user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existingUser != nil {
		return errors.New("user with this email already exists")
	}

	return s.userRepo.Create(user)
}

func (s *UserService) GetUserByID(id uint) (*entities.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*entities.User, error) {
	return s.userRepo.GetByEmail(email)
}

func (s *UserService) UpdateUser(user *entities.User) error {
	existingUser, err := s.userRepo.GetByID(user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}

	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	existingUser, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}

	return s.userRepo.Delete(id)
}

func (s *UserService) GetAllUsers(limit, offset int) ([]*entities.User, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return s.userRepo.GetAll(limit, offset)
}
