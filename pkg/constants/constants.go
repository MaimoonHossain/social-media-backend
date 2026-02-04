package constants

const (
	// User roles
	RoleUser  = "user"
	RoleAdmin = "admin"

	// Post types
	PostTypeImage = "image"
	PostTypeVideo = "video"
	PostTypeText  = "text"

	// Story duration
	StoryDuration = 24 // hours

	// Notification types
	NotificationTypeLike    = "like"
	NotificationTypeComment = "comment"
	NotificationTypeFollow  = "follow"
	NotificationTypeMention = "mention"

	// Pagination defaults
	DefaultPage     = 1
	DefaultPageSize = 20
	MaxPageSize     = 100

	// File upload
	MaxProfileImageSize = 5 * 1024 * 1024  // 5MB
	MaxPostImageSize    = 10 * 1024 * 1024 // 10MB
	MaxVideoSize        = 50 * 1024 * 1024 // 50MB

	// Cache keys
	CacheKeyUserProfile = "user:profile:%s"
	CacheKeyUserFeed    = "user:feed:%s"
	CacheKeyPost        = "post:%s"

	// Rate limiting
	RateLimitAuth   = 5   // 5 requests per minute
	RateLimitAPI    = 100 // 100 requests per minute
	RateLimitUpload = 10  // 10 uploads per minute
)

var (
	// Allowed image extensions
	AllowedImageExtensions = []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}

	// Allowed video extensions
	AllowedVideoExtensions = []string{".mp4", ".mov", ".avi", ".mkv"}
)