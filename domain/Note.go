package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	Id        string         `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	Title     string         `json:"title" gorm:"type:varchar(80);not null"`
	Body      string         `json:"body" gorm:"type:text;not null"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	Deleted   gorm.DeletedAt `json:"-"`
}

func (u *Note) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Id == "" {
		u.Id = uuid.New().String()
	}
	return
}
