package spotify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/robinho46/spotify-match.git/internal/models"
)

func GetAccessToken() (string, error) {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	if clientID == "" {
		return "", fmt.Errorf("SPOTIFY_CLIENT_ID is not set")
	}
	if clientSecret == "" {
		return "", fmt.Errorf("SPOTIFY_CLIENT_SECRET is not set")
	}

	combined := clientID + ":" + clientSecret
	data := []byte(combined)

	encoded := base64.StdEncoding.EncodeToString(data)

	url := "https://accounts.spotify.com/api/token"
	body := strings.NewReader("grant_type=client_credentials")
	req, err := http.NewRequest("POST", url, body)

	if err != nil {
		return "", fmt.Errorf("error when making POST request: %w", err)
	}
	req.Header.Set("Authorization", "Basic "+encoded)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	type tokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	var token tokenResponse
	err = json.Unmarshal(bodyBytes, &token)
	if err != nil {
		return "", fmt.Errorf("failed to decode JSON: %w", err)
	}

	return token.AccessToken, nil
}

func GetPlaylistTracks(playlistID string, token string) ([]models.Track, error) {
	const limit = 100

	page, total, err := fetchTracksPage(playlistID, token, 0, limit)

	if err != nil {
		return nil, err
	}

	result := make([]models.Track, len(page))
	copy(result, page)

	if total <= limit {
		return result, nil
	}

	var offsets []int

	for offset := 100; offset < total; offset += limit {
		offsets = append(offsets, offset)
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex
	c := make(chan int, len(offsets))
	workers := runtime.NumCPU()
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for offset := range c {
				tracks, _, err := fetchTracksPage(playlistID, token, offset, limit)
				if err != nil {
					continue
				}
				mutex.Lock()
				result = append(result, tracks...)
				mutex.Unlock()
			}
		}()
	}
	for _, offset := range offsets {
		c <- offset
	}
	close(c)
	wg.Wait()

	return result, nil
}

func fetchTracksPage(playlistID, token string, offset, limit int) ([]models.Track, int, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks?limit=%d&offset=%d", playlistID, limit, offset)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	var parsed models.PlaylistResponse
	err = json.Unmarshal(bodyBytes, &parsed)
	if err != nil {
		return nil, 0, err
	}

	var result []models.Track
	for _, i := range parsed.Items {
		if len(i.Track.Artists) > 0 {
			track := models.Track{
				Name:   i.Track.Name,
				ID:     i.Track.ID,
				Artist: i.Track.Artists[0].Name,
			}
			result = append(result, track)
		}

	}
	return result, parsed.Total, nil
}
