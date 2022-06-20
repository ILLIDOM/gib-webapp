package domain

import "time"

type Report struct {
	ID           int
	CreationDate time.Time
	Content      string
}
