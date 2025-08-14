package posts

import "time"

type CreateCommentRequest struct {
	CommentContent string `json:"commentContent"`
}

type CommentModel struct {
	Id             int64     `db:"id"`
	PostId         int64     `db:"post_id"`
	UserId         int64     `db:"user_id"`
	CommentContent string    `db:"comment_content"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	CreatedBy      string    `db:"created_by"`
	UpdatedBy      string    `db:"updated_by"`
}
