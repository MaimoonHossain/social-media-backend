package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Hashtag struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name      string    `gorm:"uniqueIndex;not null;size:100" json:"name"`
	PostCount int       `gorm:"default:0" json:"post_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationships
	Posts []Post `gorm:"many2many:post_hashtags;" json:"posts,omitempty"`
}

func (h *Hashtag) BeforeCreate(tx *gorm.DB) error {
	if h.ID == uuid.Nil {
		h.ID = uuid.New()
	}
	return nil
}

// HashtagResponse for hashtag data
type HashtagResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	PostCount int       `json:"post_count"`
	CreatedAt time.Time `json:"created_at"`
}