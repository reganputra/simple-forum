package posts

import (
	"context"
	"errors"
	"simple-forum/internal/model/posts"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *Service) UpsertUserActivity(ctx context.Context, postId, userId int64, req posts.UserActivityRequest) error {

	now := time.Now()
	model := posts.UserActivityModel{
		PostId:    postId,
		UserId:    userId,
		IsLiked:   req.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userId, 10),
		UpdatedBy: strconv.FormatInt(userId, 10),
	}
	userActivity, err := s.postRepo.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("Error getting user activity")
		return err
	}
	if userActivity == nil {
		if !req.IsLiked {
			return errors.New("user has not liked this post yet")
		}
		err = s.postRepo.CreateUserActivity(ctx, model)
	} else {
		err = s.postRepo.UpdateUserActivity(ctx, model)
	}
	if err != nil {
		log.Error().Err(err).Msg("Error creating user activity")
		return err
	}
	return nil
}
