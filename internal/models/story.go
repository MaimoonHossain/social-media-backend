package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Story struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID     uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	MediaURL   string    `gorm:"not null;size:255" json:"media_url"`
	MediaType  string    `gorm:"not null;size:20" json:"media_type"` // image, video
	Caption    string    `gorm:"type:text" json:"caption,omitempty"`
	ViewsCount int       `gorm:"default:0" json:"views_count"`
	ExpiresAt  time.Time `gorm:"not null;index" json:"expires_at"`
	CreatedAt  time.Time `json:"created_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	User  User         `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Views []StoryView  `gorm:"foreignKey:StoryID" json:"views,omitempty"`

	// Computed
	IsViewed bool `gorm:"-" json:"is_viewed,omitempty"`
}

func (s *Story) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	// Stories expire after 24 hours
	if s.ExpiresAt.IsZero() {
		s.ExpiresAt = time.Now().Add(24 * time.Hour)
	}
	return nil
}

// StoryView tracks who viewed a story
type StoryView struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	StoryID   uuid.UUID `gorm:"type:uuid;not null;index" json:"story_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	ViewedAt  time.Time `json:"viewed_at"`

	// Relationships
	Story Story `gorm:"foreignKey:StoryID" json:"story,omitempty"`
	User  User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (sv *StoryView) BeforeCreate(tx *gorm.DB) error {
	if sv.ID == uuid.Nil {
		sv.ID = uuid.New()
	}
	if sv.ViewedAt.IsZero() {
		sv.ViewedAt = time.Now()
	}
	return nil
}

// CreateStoryRequest for creating stories
type CreateStoryRequest struct {
	MediaType string `json:"media_type" binding:"required,oneof=image video"`
	Caption   string `json:"caption,omitempty" binding:"omitempty,max=500"`
}

// StoryResponse includes user info
type StoryResponse struct {
	ID         uuid.UUID    `json:"id"`
	User       UserResponse `json:"user"`
	MediaURL   string       `json:"media_url"`
	MediaType  string       `json:"media_type"`
	Caption    string       `json:"caption,omitempty"`
	ViewsCount int          `json:"views_count"`
	IsViewed   bool         `json:"is_viewed"`
	ExpiresAt  time.Time    `json:"expires_at"`
	CreatedAt  time.Time    `json:"created_at"`
}