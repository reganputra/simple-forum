package posts

import (
	"context"
	"simple-forum/internal/model/posts"

	"github.com/rs/zerolog/log"
)

func (s *Service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize
	resp, err := s.postRepo.GetAllPost(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("Error getting all post")
		return resp, err
	}
	return resp, nil
}
