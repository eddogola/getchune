package getchune

import (
	"regexp"
	"strings"
)

var videoRegexList = []*regexp.Regexp{
	regexp.MustCompile(`(?:watch\?v)(?:=)([^"%&?.\/=]{11})`),
	regexp.MustCompile(`(?:=)([^"%&?.\/=]{11})`),
	regexp.MustCompile(`([^"%&?.\/=]{11})`),
}

func extractVideoID(videoID string) (string, error) {
	if strings.Contains(videoID, "youtube") || strings.ContainsAny(videoID, "\"?%&=/") {
		for _, regex := range videoRegexList {
			if isMatch := regex.MatchString(videoID); isMatch {
				subs := regex.FindStringSubmatch(videoID)
				videoID = subs[0]
			}
		}
	}

	if strings.ContainsAny(videoID, "\"?%&=/") {
		return "", ErrInvalidCharsInVideoID
	}

	if len(videoID) < 10 {
		return "", ErrVideoIDMinLength
	}

	return videoID, nil
}
