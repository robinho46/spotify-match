package cmd

import (
	"testing"
)

func TestExtractPlaylistID(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"https://open.spotify.com/playlist/37i9dQZF1DXcBWIGoYBM5M", "37i9dQZF1DXcBWIGoYBM5M"},
		{"https://open.spotify.com/playlist/5dypIFDX2Xxd4gjFH4sCnv?si=97ca36779dfd46df", "5dypIFDX2Xxd4gjFH4sCnv"},
		{"invalid_url", ""},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := ExtractPlaylistID(test.input)
			if result != test.expected {
				t.Errorf("For input '%s', expected '%s', got '%s'", test.input, test.expected, result)
			}
		})
	}
}
