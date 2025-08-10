package membership

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"simple-forum/internal/model/membership"
	"time"
)

func (s *Service) SignUp(ctx context.Context, req *membership.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("user already exists with this email or username")
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	now := time.Now()
	userModel := &membership.UserModel{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(pw),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}
	err = s.membershipRepo.CreateUser(ctx, userModel)
	if err != nil {
		return err
	}
	return nil
}
