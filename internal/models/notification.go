package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"` // Who receives the notification
	ActorID   uuid.UUID  `gorm:"type:uuid;not null" json:"actor_id"` // Who triggered the notification
	Type      string     `gorm:"not null;size:50" json:"type"` // like, comment, follow, mention
	PostID    *uuid.UUID `gorm:"type:uuid" json:"post_id,omitempty"`
	CommentID *uuid.UUID `gorm:"type:uuid" json:"comment_id,omitempty"`
	Content   string     `gorm:"type:text" json:"content,omitempty"`
	IsRead    bool       `gorm:"default:false" json:"is_read"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	// Relationships
	User    User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Actor   User     `gorm:"foreignKey:ActorID" json:"actor,omitempty"`
	Post    *Post    `gorm:"foreignKey:PostID" json:"post,omitempty"`
	Comment *Comment `gorm:"foreignKey:CommentID" json:"comment,omitempty"`
}

func (n *Notification) BeforeCreate(tx *gorm.DB) error {
	if n.ID == uuid.Nil {
		n.ID = uuid.New()
	}
	return nil
}

// NotificationResponse includes actor info
type NotificationResponse struct {
	ID        uuid.UUID    `json:"id"`
	Actor     UserResponse `json:"actor"`
	Type      string       `json:"type"`
	PostID    *uuid.UUID   `json:"post_id,omitempty"`
	CommentID *uuid.UUID   `json:"comment_id,omitempty"`
	Content   string       `json:"content,omitempty"`
	IsRead    bool         `json:"is_read"`
	CreatedAt time.Time    `json:"created_at"`
}