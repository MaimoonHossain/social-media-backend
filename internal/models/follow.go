package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Follow struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	FollowerID  uuid.UUID `gorm:"type:uuid;not null;index" json:"follower_id"`
	FollowingID uuid.UUID `gorm:"type:uuid;not null;index" json:"following_id"`
	Status      string    `gorm:"default:'accepted';size:20" json:"status"` // pending, accepted
	CreatedAt   time.Time `json:"created_at"`

	// Relationships
	Follower  User `gorm:"foreignKey:FollowerID" json:"follower,omitempty"`
	Following User `gorm:"foreignKey:FollowingID" json:"following,omitempty"`
}

func (f *Follow) BeforeCreate(tx *gorm.DB) error {
	if f.ID == uuid.Nil {
		f.ID = uuid.New()
	}
	return nil
}

// Ensure unique follow relationships
func (Follow) TableName() string {
	return "follows"
}

// FollowResponse for follow relationships
type FollowResponse struct {
	ID          uuid.UUID    `json:"id"`
	Follower    UserResponse `json:"follower"`
	Following   UserResponse `json:"following"`
	Status      string       `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
}