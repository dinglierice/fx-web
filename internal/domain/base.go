package domain

import "time"

type Base struct {
	ID        int
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
