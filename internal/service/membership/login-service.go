package membership

import (
	"context"
	"errors"
	"simple-forum/internal/model/membership"
	"simple-forum/pkg/jwt"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(ctx context.Context, req membership.LoginRequest) (string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "")
	if err != nil {
		log.Error().Err(err).Msg("error getting user")
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := jwt.CreateToken(user.Id, user.Username, s.config.Service.SecretJWT)
	if err != nil {
		return "", err
	}
	return token, nil
}
