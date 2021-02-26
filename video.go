package getchune

import (
	"time"
)

type Video struct {
	ID          string
	Title       string
	Description string
	Author      string
	Duration    time.Duration
	DASHManifestURL string
	HLSManifestURL string
}
