package entity

import "time"

type SuccessReponse[T any] struct {
	StatusCode int       `json:"status_code"`
	Time       time.Time `json:"time"`
	Data       T         `json:"data"`
}
