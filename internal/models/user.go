package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Username			string     `gorm:"uniqueIndex;not null;size:50" json:"username"`
	Email					 string     `gorm:"uniqueIndex;not null;size:100" json:"email"`
	Password        string     `gorm:"not null" json:"-"`
	FullName        string     `gorm:"size:100" json:"full_name"`
	Bio             string     `gorm:"type:text" json:"bio"`
	ProfileImageURL string     `gorm:"size:255" json:"profile_image_url"`
	CoverImageURL   string     `gorm:"size:255" json:"cover_image_url"`
	Website         string     `gorm:"size:100" json:"website"`
	Location        string     `gorm:"size:100" json:"location"`
	DateOfBirth     *time.Time `json:"date_of_birth"`
	Role            string     `gorm:"default:'user';size:20" json:"role"`
	IsVerified      bool       `gorm:"default:false" json:"is_verified"`
	IsPrivate       bool       `gorm:"default:false" json:"is_private"`
	IsActive        bool       `gorm:"default:true" json:"is_active"`
	LastLoginAt     *time.Time `json:"last_login_at"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

		// Relationships
	Posts         []Post         `gorm:"foreignKey:UserID" json:"posts,omitempty"`
	Comments      []Comment      `gorm:"foreignKey:UserID" json:"comments,omitempty"`
	Likes         []Like         `gorm:"foreignKey:UserID" json:"likes,omitempty"`
	Stories       []Story        `gorm:"foreignKey:UserID" json:"stories,omitempty"`
	Followers     []Follow       `gorm:"foreignKey:FollowingID" json:"followers,omitempty"`
	Following     []Follow       `gorm:"foreignKey:FollowerID" json:"following,omitempty"`
	SentMessages  []Message      `gorm:"foreignKey:SenderID" json:"sent_messages,omitempty"`
	ReceivedMessages []Message   `gorm:"foreignKey:ReceiverID" json:"received_messages,omitempty"`
	Notifications []Notification `gorm:"foreignKey:UserID" json:"notifications,omitempty"`

	// Counts (not stored in DB, computed)
	FollowersCount int `gorm:"-" json:"followers_count,omitempty"`
	FollowingCount int `gorm:"-" json:"following_count,omitempty"`
	PostsCount     int `gorm:"-" json:"posts_count,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

// UserResponse is used for public user data
type UserResponse struct {
	ID              uuid.UUID  `json:"id"`
	Username        string     `json:"username"`
	Email           string     `json:"email,omitempty"` // Only shown to self
	FullName        string     `json:"full_name"`
	Bio             string     `json:"bio"`
	ProfileImageURL string     `json:"profile_image_url"`
	CoverImageURL   string     `json:"cover_image_url"`
	Website         string     `json:"website"`
	Location        string     `json:"location"`
	IsVerified      bool       `json:"is_verified"`
	IsPrivate       bool       `json:"is_private"`
	FollowersCount  int        `json:"followers_count"`
	FollowingCount  int        `json:"following_count"`
	PostsCount      int        `json:"posts_count"`
	IsFollowing     bool       `json:"is_following,omitempty"`
	IsFollowedBy    bool       `json:"is_followed_by,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
}

// LoginRequest for user login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterRequest for user registration
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required,min=1,max=100"`
}

// UpdateProfileRequest for updating user profile
type UpdateProfileRequest struct {
	FullName  string `json:"full_name,omitempty" binding:"omitempty,max=100"`
	Bio       string `json:"bio,omitempty" binding:"omitempty,max=500"`
	Website   string `json:"website,omitempty" binding:"omitempty,url,max=100"`
	Location  string `json:"location,omitempty" binding:"omitempty,max=100"`
	IsPrivate *bool  `json:"is_private,omitempty"`
}