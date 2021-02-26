package getchune

import (
	"context"
	"io/ioutil"
	"net/http"
)

// Client serves requests to get YT videos
type Client struct {
	HTTPClient *http.Client
}

// GetVideo gets a Video object from a YT video ID
func (c *Client) GetVideo(videoID string) (*Video, error) {
	return c.GetVideoWithContext(context.Background(), videoID)
}

// GetVideoWithContext gets a Video object with a context
func (c *Client) GetVideoWithContext(ctx context.Context, videoID string) (*Video, error) {
	id, err := extractVideoID(videoID)
	if err != nil {
		return nil, err
	}

	// circumvent age restriction; pretend to come from googleapis
	eurl := "https://youtube.googleapis.com/v/" + id
	body, err := c.httpGetBodyBytes(ctx, "https://youtube.com/get_video_info?video_id="+id+"&eurl="+eurl)
	if err != nil {
		return nil, err
	}

	v := &Video{
		ID: id,
	}

	err = v.extractVideoInfo(body)

	return v, err
}

// GetStream returns the http.Response for a specific format
func (c *Client) GetStream(v *Video, format *Format) (*http.Response, error) {
	return c.GetStreamWithContext(context.Background(), v, format)
}

// GetStreamWithContext returns the http.Response for a specific format with a context
func (c *Client) GetStreamWithContext(ctx context.Context, v *Video, format *Format) (*http.Response, error) {
	url, err := c.GetStreamURL(ctx, v, format)
	if err != nil {
		return nil, err
	}

	return c.httpGet(ctx, url)
}

// GetStreamURL returns the url of a specific format
func (c *Client) GetStreamURL(ctx context.Context, v *Video, format *Format) (string, error) {
	return c.GetStreamURLWithContext(ctx, v, format)
}

// GetStreamURLWithContext returns the url of a specific format with a context
func (c *Client) GetStreamURLWithContext(ctx context.Context, v *Video, format *Format) (string, error) {
	if format.URL != "" {
		return format.URL, nil
	}

	cipher := format.Cipher
	if cipher == "" {
		return "", ErrCipherNotFound
	}

	return c.decipherURL(ctx, v.ID, format)
}


func (c *Client) httpGet(ctx context.Context, url string) (*http.Response, error) {
	client := c.HTTPClient
	if client == nil {
		c.HTTPClient = http.DefaultClient
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrUnexpectedStatusCode(resp.StatusCode)
	}

	return resp, err
}

func (c *Client) httpGetBodyBytes(ctx context.Context, url string) ([]byte, error) {
	resp, err := c.httpGet(ctx, url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
