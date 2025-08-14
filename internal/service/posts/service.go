package posts

import (
	"context"
	"simple-forum/internal/configs"
	"simple-forum/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error
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
