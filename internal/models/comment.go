package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	PostID     uuid.UUID `gorm:"type:uuid;not null;index" json:"post_id"`
	UserID     uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	ParentID   *uuid.UUID `gorm:"type:uuid;index" json:"parent_id,omitempty"` // For nested comments
	Content    string    `gorm:"type:text;not null" json:"content"`
	LikesCount int       `gorm:"default:0" json:"likes_count"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Post Post `gorm:"foreignKey:PostID" json:"post,omitempty"`
	Parent *Comment `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Replies []Comment `gorm:"foreignKey:ParentID" json:"replies,omitempty"`
	Likes []Like `gorm:"foreignKey:CommentID" json:"likes,omitempty"`

	// Computed fields
	IsLiked bool `gorm:"-" json:"is_liked,omitempty"`
	RepliesCount int `gorm:"-" json:"replies_count,omitempty"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

// CreateCommentRequest for creating a comment
type CreateCommentRequest struct {
	PostID   uuid.UUID  `json:"post_id" binding:"required"`
	ParentID *uuid.UUID `json:"parent_id,omitempty"`
	Content  string     `json:"content" binding:"required,min=1,max=1000"`
}

// UpdateCommentRequest for updating a comment
type UpdateCommentRequest struct {
	Content string `json:"content" binding:"required,min=1,max=1000"`
}

// CommentResponse includes user info
type CommentResponse struct {
	ID           uuid.UUID    `json:"id"`
	PostID       uuid.UUID    `json:"post_id"`
	User         UserResponse `json:"user"`
	ParentID     *uuid.UUID   `json:"parent_id,omitempty"`
	Content      string       `json:"content"`
	LikesCount   int          `json:"likes_count"`
	RepliesCount int          `json:"replies_count"`
	IsLiked      bool         `json:"is_liked"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}