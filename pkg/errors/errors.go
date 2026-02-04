package errors

import "errors"

var (
	// Authentication errors
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrTokenExpired       = errors.New("token expired")
	ErrInvalidToken       = errors.New("invalid token")

	// User errors
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrEmailAlreadyUsed  = errors.New("email already in use")
	ErrUsernameAlreadyUsed = errors.New("username already in use")

	// Post errors
	ErrPostNotFound   = errors.New("post not found")
	ErrUnauthorizedAction = errors.New("unauthorized to perform this action")

	// Comment errors
	ErrCommentNotFound = errors.New("comment not found")

	// Follow errors
	ErrAlreadyFollowing = errors.New("already following this user")
	ErrNotFollowing     = errors.New("not following this user")
	ErrCannotFollowSelf = errors.New("cannot follow yourself")

	// Like errors
	ErrAlreadyLiked = errors.New("already liked")
	ErrNotLiked     = errors.New("not liked yet")

	// Message errors
	ErrMessageNotFound = errors.New("message not found")
	ErrCannotMessageSelf = errors.New("cannot message yourself")

	// Story errors
	ErrStoryNotFound = errors.New("story not found")
	ErrStoryExpired  = errors.New("story has expired")

	// File errors
	ErrInvalidFileType = errors.New("invalid file type")
	ErrFileTooLarge    = errors.New("file size exceeds limit")
	ErrFileUploadFailed = errors.New("file upload failed")

	// Validation errors
	ErrInvalidInput = errors.New("invalid input")
	ErrValidationFailed = errors.New("validation failed")

	// General errors
	ErrInternalServer = errors.New("internal server error")
	ErrNotFound       = errors.New("resource not found")
	ErrBadRequest     = errors.New("bad request")
)