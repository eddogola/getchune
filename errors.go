package getchune

import (
	"errors"
)

var (
	// ErrInvalidCharsInVideoID defines error when video ID has any of #%&</"
	ErrInvalidCharsInVideoID = errors.New("invalid characters in video ID")
	// ErrVideoIDMinLength defines error when video ID's length is less than 10
	ErrVideoIDMinLength      = errors.New("Video ID less than min length")
)
