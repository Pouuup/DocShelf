package models

import "time"

type Document struct {
	Title      string
	SourceURL  string
	Sections   []Section
	ImportedAt time.Time
	UpdatedAt  time.Time
}
