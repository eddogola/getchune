package getchune

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

// Video offers methods to get video metadata.
type Video struct {
	ID              string
	Title           string
	Description     string
	Formats         []Format
	Author          string
	Duration        time.Duration
	DASHManifestURL string
	HLSManifestURL  string
}

func (v *Video) extractVideoInfo(data []byte) error {
	answer, err := url.ParseQuery(string(data))
	if err != nil {
		return err
	}

	status := answer.Get("status")
	if status != "ok" {
		return ErrResponseStatus{
			status: status,
			reason: answer.Get("reason"),
		}
	}

	playerResponse := answer.Get("player_response")
	if playerResponse == "" {
		return ErrPlayerResponseNotFoundInAnswer
	}
	var prData playerResponseData

	if err = json.Unmarshal([]byte(playerResponse), &prData); err != nil {
		return err
	}

	if err = v.isVideoFromInfoDownloadable(prData); err != nil {
		return nil
	}
	return v.extractDataFromPlayerResponse(prData)
}

func (v *Video) isVideoFromInfoDownloadable(prData playerResponseData) error {
	return v.isVideoDownloadable(prData)
}

func (v *Video) isVideoDownloadable(prData playerResponseData) error {
	if prData.PlayabilityStatus.Status != "ok" {
		return ErrPlayabilityStatusNotOk
	}
	return nil
}

func (v *Video) extractDataFromPlayerResponse(prData playerResponseData) error {
	v.Title = prData.VideoDetails.Title
	v.Author = prData.VideoDetails.Author
	v.Description = prData.VideoDetails.Description
	secs, err := strconv.Atoi(prData.VideoDetails.LengthSeconds)
	if err != nil {
		return err
	}
	v.Duration = time.Duration(secs * int(time.Second))
	v.DASHManifestURL = prData.StreamingData.DashManifestURL
	v.HLSManifestURL = prData.StreamingData.HlsManifestURL
	v.Formats = append(prData.StreamingData.Formats, prData.StreamingData.AdaptiveFormats...)

	if len(v.Formats) == 0 {
		return ErrNoFormats
	}

	return nil
}
