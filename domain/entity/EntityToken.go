package entity

type EntityToken struct {
	Token     string `json:"token"`
	CreatedAt int64  `json:"created_at"`
	ExpiredIn int    `json:"expired_in"`
}
