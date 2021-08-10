package app

import "time"

type Slot struct {
	Id int32 `json:"id"`
	Date time.Time `json:"date"`
}
