package applayer

import "context"

type Post struct {
	Content string
	Owner   string
}

func (a *app) CreatePost(ctx context.Context, content, owner string) error {
	return a.store.CreatePost(ctx, content, owner)
}

func (a *app) GetAllPosts(ctx context.Context) ([]Post, error) {
	postRows, err := a.store.GetAllPosts(ctx)
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, postRow := range postRows {
		posts = append(posts, Post{
			Owner:   postRow.User.Handle,
			Content: postRow.Content,
		})
	}
	return posts, nil
}
