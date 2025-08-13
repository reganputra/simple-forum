package membership

import (
	"context"
	"simple-forum/internal/configs"
	"simple-forum/internal/model/membership"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string) (*membership.UserModel, error)
	CreateUser(ctx context.Context, user *membership.UserModel) error
}

type Service struct {
	config         *configs.Config
	membershipRepo membershipRepository
}

func NewService(cfg *configs.Config, membershipRepo membershipRepository) *Service {
	return &Service{
		config:         cfg,
		membershipRepo: membershipRepo,
	}
}
