package getchune

import (
	"testing"
)

func TestExtractVideoID(t *testing.T) {
	tests := []struct {
		name        string
		link        string
		got         string
		want        string
		wantErr     bool
		expectedErr error
	}{
		{
			name:        "valid url",
			link:        "https://www.youtube.com/watch?v=xisZPyQGsKw",
			want:        "xisZPyQGsKw",
			wantErr:     false,
			expectedErr: nil,
		},
		{
			name:        "valid video ID",
			link:        "xisZPyQGsKw",
			want:        "xisZPyQGsKw",
			wantErr:     false,
			expectedErr: nil,
		},
		{
			name:        "invalid url",
			link:        "https://www.youtube.com/watch?v=xisZP/yQG/sKw",
			want:        "",
			wantErr:     true,
			expectedErr: ErrInvalidCharsInVideoID,
		},
		{
			name:        "video ID small length",
			link:        "xisZPyKw",
			want:        "",
			wantErr:     true,
			expectedErr: ErrVideoIDMinLength,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractVideoID(tt.link)
			if tt.wantErr && (tt.expectedErr != err || err == nil) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			}
			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
