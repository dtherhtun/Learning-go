package storelayer

import (
	"context"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID  uint
	Content string
	User    User
}

func (s *store) CreatePost(ctx context.Context, content, owner string) error {
	user := &User{}
	err := s.db.Limit(1).Find(user, "handle = ?", owner).Error
	if err != nil {
		return err
	}

	err = s.db.WithContext(ctx).Create(&Post{
		Content: content,
		UserID:  user.ID,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *store) GetAllPosts(ctx context.Context) ([]Post, error) {
	var posts []Post

	err := s.db.WithContext(ctx).Preload("User").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
