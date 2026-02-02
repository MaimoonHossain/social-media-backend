package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	SenderID   uuid.UUID `gorm:"type:uuid;not null;index" json:"sender_id"`
	ReceiverID uuid.UUID `gorm:"type:uuid;not null;index" json:"receiver_id"`
	Content    string    `gorm:"type:text;not null" json:"content"`
	MediaURL   string    `gorm:"size:255" json:"media_url,omitempty"`
	MediaType  string    `gorm:"size:20" json:"media_type,omitempty"` // image, video
	IsRead     bool      `gorm:"default:false" json:"is_read"`
	ReadAt     *time.Time `json:"read_at,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Sender   User `gorm:"foreignKey:SenderID" json:"sender,omitempty"`
	Receiver User `gorm:"foreignKey:ReceiverID" json:"receiver,omitempty"`
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

// SendMessageRequest for sending messages
type SendMessageRequest struct {
	ReceiverID uuid.UUID `json:"receiver_id" binding:"required"`
	Content    string    `json:"content" binding:"required,min=1,max=5000"`
	MediaType  string    `json:"media_type,omitempty" binding:"omitempty,oneof=image video"`
}

// MessageResponse includes sender and receiver info
type MessageResponse struct {
	ID         uuid.UUID    `json:"id"`
	Sender     UserResponse `json:"sender"`
	Receiver   UserResponse `json:"receiver"`
	Content    string       `json:"content"`
	MediaURL   string       `json:"media_url,omitempty"`
	MediaType  string       `json:"media_type,omitempty"`
	IsRead     bool         `json:"is_read"`
	ReadAt     *time.Time   `json:"read_at,omitempty"`
	CreatedAt  time.Time    `json:"created_at"`
}

// Conversation represents a conversation between two users
type Conversation struct {
	User          UserResponse `json:"user"`
	LastMessage   MessageResponse `json:"last_message"`
	UnreadCount   int          `json:"unread_count"`
}