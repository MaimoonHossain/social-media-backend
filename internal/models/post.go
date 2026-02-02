package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID        uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	Caption       string         `gorm:"type:text" json:"caption"`
	MediaURL      string         `gorm:"size:255" json:"media_url"`
	MediaType     string         `gorm:"size:20" json:"media_type"` // image, video, text
	LikesCount    int            `gorm:"default:0" json:"likes_count"`
	CommentsCount int            `gorm:"default:0" json:"comments_count"`
	SharesCount   int            `gorm:"default:0" json:"shares_count"`
	ViewsCount    int            `gorm:"default:0" json:"views_count"`
	IsPublic      bool           `gorm:"default:true" json:"is_public"`
	Location      string         `gorm:"size:100" json:"location,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	User     User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
	Likes    []Like    `gorm:"foreignKey:PostID" json:"likes,omitempty"`
	Hashtags []Hashtag `gorm:"many2many:post_hashtags;" json:"hashtags,omitempty"`

	// Computed fields
	IsLiked bool `gorm:"-" json:"is_liked,omitempty"`
	IsSaved bool `gorm:"-" json:"is_saved,omitempty"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

// CreatePostRequest for creating a new post
type CreatePostRequest struct {
	Caption   string `json:"caption" binding:"omitempty,max=2200"`
	MediaType string `json:"media_type" binding:"omitempty,oneof=image video text"`
	Location  string `json:"location" binding:"omitempty,max=100"`
	IsPublic  *bool  `json:"is_public"`
}

// UpdatePostRequest for updating a post
type UpdatePostRequest struct {
	Caption  string `json:"caption" binding:"omitempty,max=2200"`
	Location string `json:"location" binding:"omitempty,max=100"`
	IsPublic *bool  `json:"is_public"`
}

// PostResponse includes user and engagement info
type PostResponse struct {
	ID            uuid.UUID    `json:"id"`
	User          UserResponse `json:"user"`
	Caption       string       `json:"caption"`
	MediaURL      string       `json:"media_url"`
	MediaType     string       `json:"media_type"`
	LikesCount    int          `json:"likes_count"`
	CommentsCount int          `json:"comments_count"`
	SharesCount   int          `json:"shares_count"`
	ViewsCount    int          `json:"views_count"`
	Location      string       `json:"location,omitempty"`
	IsLiked       bool         `json:"is_liked"`
	IsSaved       bool         `json:"is_saved"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}