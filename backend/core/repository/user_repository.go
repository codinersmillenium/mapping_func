package repository

import "mapping_func/core/entity"

type UserRepository interface {
	Insert(user *entity.User) error
}

type userRepository struct {
	db DB
}

func NewUserRepository(db DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Insert(user *entity.User) error {
	q := `INSERT INTO users(name, age, city) VALUES (?, ?, ?)`

	_, err := r.db.Exec(
		r.db.Rebind(q),
		user.Name,
		user.Age,
		user.City,
	)

	return err
}
