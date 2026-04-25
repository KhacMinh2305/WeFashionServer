package entity

type EntityToken struct {
	Token     string `json:"token"`
	CreatedAt int64  `josn:"created_at"`
	ExpiredIn int    `josn:"expired_in"`
}
