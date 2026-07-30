package models

import "time"

type Document struct {
	Title       string
	Source_URL  string
	Description string
	Imported_At time.Time
	Updated_At  time.Time
}
