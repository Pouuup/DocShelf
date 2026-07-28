package models

import "time"

type Document struct {
	Title       string
	SourceURL   string
	Description string
	ImportedAt  time.Time
	UpdatedAt   time.Time
}
