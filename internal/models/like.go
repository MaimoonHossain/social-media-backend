package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Like struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	PostID    *uuid.UUID `gorm:"type:uuid;index" json:"post_id,omitempty"`
	CommentID *uuid.UUID `gorm:"type:uuid;index" json:"comment_id,omitempty"`
	CreatedAt time.Time  `json:"created_at"`

	// Relationships
	User    User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Post    *Post    `gorm:"foreignKey:PostID" json:"post,omitempty"`
	Comment *Comment `gorm:"foreignKey:CommentID" json:"comment,omitempty"`
}

func (l *Like) BeforeCreate(tx *gorm.DB) error {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return nil
}

// LikeRequest for liking posts or comments
type LikeRequest struct {
	PostID    *uuid.UUID `json:"post_id,omitempty"`
	CommentID *uuid.UUID `json:"comment_id,omitempty"`
}