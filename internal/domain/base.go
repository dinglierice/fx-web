package domain

import "time"

type Base struct {
	ID        uint64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
