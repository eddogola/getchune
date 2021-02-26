package getchune

import (
	"errors"
)

var (
	ErrInvalidCharsInVideoID = errors.New("invalid characters in video ID")
	ErrVideoIDMinLength      = errors.New("Video ID less than min length")
)
