package models

import "time"

type Problem struct {
	ID int64
	Slug string
	Title string
	Description string
	Difficulty string
	TimeLimitMs int
	MemoryLimitMb int
	CreatedAt time.Time
}

type CreateProblemRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Difficulty string `json:"difficulty"`
}
