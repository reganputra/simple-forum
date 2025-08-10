package membership

import (
	"context"
	"simple-forum/internal/model/membership"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string) (*membership.UserModel, error)
	CreateUser(ctx context.Context, user *membership.UserModel) error
}

type Service struct {
	membershipRepo membershipRepository
}

func NewService(membershipRepo membershipRepository) *Service {
	return &Service{
		membershipRepo: membershipRepo,
	}
}
