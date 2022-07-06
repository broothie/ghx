package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	gorm.Model
	ID string `gorm:"type:uuid;"`
}

func ID(id string) Base {
	return Base{ID: id}
}

func (b *Base) BeforeCreate(*gorm.DB) error {
	b.ID = uuid.NewString()
	return nil
}

type Post struct {
	Base
	Title    string
	Link     string
	Votes    int
	Comments []Comment
}

type Comment struct {
	Base
	Text   string
	Votes  int
	PostID string
	Post   Post
}
