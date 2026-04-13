package entity

import "time"

type EntityToken struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `josn:"created_at"`
	ExpiredIn int       `josn:"expired_in"`
}
