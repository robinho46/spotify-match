package spotify

import (
	"github.com/robinho46/spotify-match.git/internal/models"
	"testing"
)

func TestComparePlaylists(t *testing.T) {
	tracks1 := []models.Track{
		{Name: "Song A", ID: "1", Artist: "Artist X"},
		{Name: "Song B", ID: "2", Artist: "Artist Y"},
	}

	tracks2 := []models.Track{
		{Name: "Song A", ID: "1", Artist: "Artist X"}, // Same song
		{Name: "Song C", ID: "3", Artist: "Artist Y"}, // Same artist
	}

	commonTracks, commonArtists, score := ComparePlaylists(tracks1, tracks2)

	if commonTracks != 1 {
		t.Errorf("Expected 1 common track, got %d", commonTracks)
	}
	if commonArtists != 2 {
		t.Errorf("Expected 2 common artists, got %d", commonArtists)
	}
	if score <= 0 {
		t.Errorf("Expected positive similarity score, got %f", score)
	}
}
