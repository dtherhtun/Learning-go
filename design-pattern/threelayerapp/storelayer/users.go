package storelayer

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Handle string
	Name   string
	Posts  []Post
}

func (s *store) CreateUser(ctx context.Context, name, handle string) error {
	err := s.db.WithContext(ctx).Create(&User{
		Name:   name,
		Handle: handle,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *store) GetAllUsers(ctx context.Context) ([]User, error) {
	var users []User
	err := s.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
