package service

import (
	"fmt"
	"mapping_func/core/entity"
	"mapping_func/core/repository"
	"strings"
)

type UserService struct {
	repo   repository.UserRepository
	cities map[string]string
}

func NewUserService(r repository.UserRepository, c map[string]string) *UserService {
	return &UserService{r, c}
}

func (s *UserService) Create(input string) error {
	parts := strings.Split(input, " ")

	if len(parts) < 3 {
		return nil
	}

	age := parts[len(parts)-2]
	city := strings.ToUpper(parts[len(parts)-1])
	name := strings.ToUpper(strings.Join(parts[:len(parts)-2], " "))
	fmt.Println(age)
	if p, ok := s.cities[city]; ok {
		city = city + " " + p
	}

	u := &entity.User{
		Name: name,
		Age:  age,
		City: city,
	}

	return s.repo.Insert(u)
}
