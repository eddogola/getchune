package getchune

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidCharsInVideoID defines error when video ID has any of #%&</"
	ErrInvalidCharsInVideoID = errors.New("invalid characters in video ID")
	// ErrVideoIDMinLength defines error when video ID's length is less than 10
	ErrVideoIDMinLength = errors.New("Video ID less than min length")
	// ErrPlayerResponseNotFoundInAnswer when player_response is ""
	ErrPlayerResponseNotFoundInAnswer = errors.New("player_response not found in answer")
	// ErrNoFormats no formats found
	ErrNoFormats = errors.New("no new formats found in server")
	// ErrPlayabilityStatusNotOk video can't be downloaded
	ErrPlayabilityStatusNotOk = errors.New("video playability status not ok")
	// ErrCipherNotFound cipher for specific video not found
	ErrCipherNotFound = errors.New("cipher not found")
)

// ErrUnexpectedStatusCode associated with http status code
type ErrUnexpectedStatusCode int

func (err ErrUnexpectedStatusCode) Error() string {
	return fmt.Sprintf("unexpected status code: %v", err)
}

// ErrResponseStatus : error on getting player response data
type ErrResponseStatus struct {
	status string
	reason string
}

func (err ErrResponseStatus) Error() string {
	return fmt.Sprintf("cannot playback and download, status: %v, reason: %v", err.status, err.reason)
}
