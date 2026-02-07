package service

import (
	"mapping_func/core/adapter"
	"mapping_func/core/entity"
	"mapping_func/core/repository"
)

type UserService struct {
	repo   repository.UserRepository
	parser *adapter.ParserAdapter
}

func NewUserService(r repository.UserRepository, cities map[string]string) *UserService {
	parser := adapter.NewParserAdapter(cities)
	return &UserService{
		repo:   r,
		parser: parser,
	}
}

func (s *UserService) Create(input string) error {
	name, age, city := s.parser.Parse(input)
	u := &entity.User{
		Name: name,
		Age:  age,
		City: city,
	}
	return s.repo.Insert(u)
}
