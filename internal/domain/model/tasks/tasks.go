package model

import (
	"time"
)

type Tasks struct {
	ID          int64
	Title       string
	Done        bool
	InputObject InputObject `gorm:"type:jsonb"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DoneAt      *time.Time `gorm:"default:null"`
}

type InputObject struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
	// Adicione outros campos conforme necessário
}
