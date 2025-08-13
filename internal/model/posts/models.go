package posts

import "time"

type CreatePostRequest struct {
	PostTitle     string   `json:"postTitle"`
	PostContent   string   `json:"postContent"`
	PostsHashtags []string `json:"postsHashtags"`
}

type PostModel struct {
	Id            int64     `db:"id"`
	UserId        int64     `db:"user_Id"`
	PostTitle     string    `db:"postTitle"`
	PostContent   string    `db:"postContent"`
	PostsHashtags string    `db:"postsHashtags"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
	CreatedBy     string    `db:"created_by"`
	UpdatedBy     string    `db:"updated_by"`
}
