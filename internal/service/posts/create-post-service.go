package posts

import (
	"context"
	"simple-forum/internal/model/posts"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *Service) CreatePost(ctx context.Context, userId int64, req posts.CreatePostRequest) error {
	postHastags := strings.Join(req.PostsHashtags, ",")

	now := time.Now()
	model := posts.PostModel{
		UserId:        userId,
		PostTitle:     req.PostTitle,
		PostContent:   req.PostContent,
		PostsHashtags: postHastags,
		CreatedAt:     now,
		UpdatedAt:     now,
		CreatedBy:     strconv.FormatInt(userId, 10),
		UpdatedBy:     strconv.FormatInt(userId, 10),
	}
	err := s.postRepo.CreatePost(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("Error creating post to repository")
		return err
	}
	return nil
}
