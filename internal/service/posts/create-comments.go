package posts

import (
	"context"
	"simple-forum/internal/model/posts"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *Service) CreateComment(ctx context.Context, postId, userId int64, req posts.CreateCommentRequest) error {

	now := time.Now()
	model := posts.CommentModel{
		PostId:         postId,
		UserId:         userId,
		CommentContent: req.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      strconv.FormatInt(userId, 10),
		UpdatedBy:      strconv.FormatInt(userId, 10),
	}

	err := s.postRepo.CreateComment(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("Error creating comment to repository")
		return err
	}
	return nil
}
