package model

type Follower struct {
	UserID     int64  `json:"user_id"`
	FollowerID int64  `json:"user_id"`
	CreatedAt  string `json:"created_at"`
}
