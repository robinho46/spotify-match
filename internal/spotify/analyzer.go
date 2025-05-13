package spotify

import (
	"github.com/robinho46/spotify-match.git/internal/models"
)

func ComparePlaylists(tracks1, tracks2 []models.Track) (commonTracks int, commonArtists int, similarityScore float64) {
	trackIDs := make(map[string]bool)
	for _, t := range tracks1 {
		trackIDs[t.ID] = true
	}

	for _, t := range tracks2 {
		if trackIDs[t.ID] {
			commonTracks++
		}
	}

	artistNames := make(map[string]bool)
	for _, t := range tracks1 {
		artistNames[t.Artist] = true
	}
	seen := make(map[string]bool)

	for _, t := range tracks2 {
		if artistNames[t.Artist] && !seen[t.Artist] {
			commonArtists++
			seen[t.Artist] = true
		}
	}

	similarityScore = float64(commonTracks) / float64(min(len(tracks1), len(tracks2))) * 100

	return commonTracks, commonArtists, similarityScore
}
