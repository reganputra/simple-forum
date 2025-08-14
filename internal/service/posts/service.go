package posts

import (
	"context"
	"simple-forum/internal/configs"
	"simple-forum/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error
	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error
}

type Service struct {
	config   *configs.Config
	postRepo postRepository
}

func NewService(cfg *configs.Config, postRepo postRepository) *Service {
	return &Service{
		config:   cfg,
		postRepo: postRepo,
	}
}
