package posts

import (
	"context"
	"simple-forum/internal/model/posts"
	"strings"
)

func (r *Repository) CreatePost(ctx context.Context, model posts.PostModel) error {
	query := "INSERT INTO posts (user_id, post_title, post_content, post_hashtags, created_at, updated_at, created_by, updated_by) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := r.db.ExecContext(ctx, query, model.UserId, model.PostTitle, model.PostContent, model.PostsHashtags, model.CreatedAt,
		model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostResponse, error) {

	query := `
        SELECT 
            p.id, 
            p.user_id, 
            u.username, 
            p.post_title, 
            p.post_content, 
            p.post_hashtags
        FROM posts p
        JOIN users u ON p.user_id = u.id
        ORDER BY p.updated_at DESC
        LIMIT ? OFFSET ?`

	resp := posts.GetAllPostResponse{}
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return posts.GetAllPostResponse{}, err
	}
	defer rows.Close()

	data := make([]posts.Post, 0)
	for rows.Next() {
		var model posts.PostModel
		var username string
		err = rows.Scan(&model.Id, &model.UserId, &username, &model.PostTitle, &model.PostContent, &model.PostsHashtags)
		if err != nil {
			return resp, err
		}
		data = append(data, posts.Post{
			Id:            model.Id,
			UserId:        model.UserId,
			Username:      username,
			PostTitle:     model.PostTitle,
			PostContent:   model.PostContent,
			PostsHashtags: strings.Split(model.PostsHashtags, ","),
		})
	}
	resp.Data = data
	resp.Pagination = posts.Paginataion{
		Limit:  limit,
		Offset: offset,
	}
	return resp, nil
}
