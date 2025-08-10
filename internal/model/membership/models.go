package membership

import "time"

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserModel struct {
	Id        int64     `db:"id"`
	Email     string    `db:"email"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedBy string    `db:"created_by"`
	UpdatedBy string    `db:"updated_by"`
}
