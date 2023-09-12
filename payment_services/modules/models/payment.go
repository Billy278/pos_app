package models

import "time"

type Payment struct {
	Id         uint64
	Name       string
	Type       string
	Created_At *time.Time
	Updated_At *time.Time
}
