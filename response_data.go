package getchune

type playerResponseData struct {
	PlayabilityStatus struct {
		Status          string `json:"status"`
		Reason          string `json:"reason"`
		PlayableInEmbed string `json:"playableInEmbed"`
		ContextParams   string `json:"contextParams"`
	} `json:"playabilityStatus"`
	StreamingData struct {
		ExpiresInSeconds string   `json:"expiresInSeconds"`
		Formats          []Format `json:"formats"`
		AdaptiveFormats  []Format `json:"adaptiveFormats"`
		DashManifestURL  string   `json:"dashManifestUrl"`
		HlsManifestURL   string   `json:"hlsManifestUrl"`
	} `json:"streamingData"`
	VideoDetails struct {
		Title         string `json:"title"`
		LengthSeconds string `json:"lengthSeconds"`
		ChannelID     string `json:"channelId"`
		Description   string `json:"shortDescription"`
		Author        string `json:"author"`
	}
}

// Format : video formats
type Format struct {
	ItagNo           int    `json:"itag"`
	URL              string `json:"url"`
	MimeType         string `json:"mimeType"`
	Bitrate          int    `json:"bitrate"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	Quality          string `json:"quality"`
	QualityLabel     string `json:"qualityLabel"`
	FPS              int    `json:"fps"`
	AverageBitrate   int    `json:"averageBitrate"`
	ContentLength    string `json:"contentLength"`
	ProjectionType   string `json:"projectionType"`
	ApproxDurationMs string `json:"approxDurationMs"`
	AudioQuality     string `json:"audioQuality"`
	AudioSample      string `json:"audioSampleRate"`
	AudioChannels    int    `json:"audioChannels"`

	InitRange struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"initRange"`
	IndexRange struct {
		Start string `json:"start"`
		End   string `json:"end"`
	}
}
