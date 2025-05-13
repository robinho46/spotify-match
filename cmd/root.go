package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/robinho46/spotify-match.git/internal/spotify"
)

func ExtractPlaylistID(link string) string {
	link = strings.TrimSpace(link)
	if strings.Contains(link, "playlist/") {
		parts := strings.Split(link, "playlist/")
		if len(parts) > 1 {
			id := strings.Split(parts[1], "?")[0]
			return id
		}
	}
	return ""
}

func spinner(done <-chan bool, text string) {
	symbols := []string{"|", "/", "-", "\\"}
	i := 0
	for {
		select {
		case <-done:
			fmt.Printf("\r\033[K\033[32mDone!\033[0m\n")
			return
		default:
			fmt.Printf("\r%s %s", text, symbols[i%len(symbols)])
			time.Sleep(100 * time.Millisecond)
			i++
		}
	}
}

func Execute() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	token, err := spotify.GetAccessToken()
	if err != nil {
		log.Fatal("failed to get access token:", err)
	}

	reader := bufio.NewReader(os.Stdin)

	var link1, link2 string
	for {
		fmt.Print("Paste first playlist link:\n> ")
		link1, _ = reader.ReadString('\n')
		link1 = strings.TrimSpace(link1)
		if strings.HasPrefix(link1, "https://open.spotify.com/playlist/") {
			break
		}
		fmt.Println("Invalid link")
	}
	for {
		fmt.Print("Paste second playlist link:\n> ")
		link2, _ = reader.ReadString('\n')
		link2 = strings.TrimSpace(link2)
		if strings.HasPrefix(link2, "https://open.spotify.com/playlist/") {
			break
		}
		fmt.Println("Invalid link")
	}

	done := make(chan bool)
	go spinner(done, "Analyzing playlists")

	id1 := ExtractPlaylistID(link1)
	id2 := ExtractPlaylistID(link2)

	tracks1, err := spotify.GetPlaylistTracks(id1, token)
	if err != nil {
		log.Fatal("failed to get playlist 1:", err)
	}
	tracks2, err := spotify.GetPlaylistTracks(id2, token)
	if err != nil {
		log.Fatal("failed to get playlist 2:", err)
	}
	done <- true

	commonTracks, commonArtists, score := spotify.ComparePlaylists(tracks1, tracks2)

	fmt.Println("\nSummary:")
	fmt.Printf("   Common songs:    %d\n", commonTracks)
	fmt.Printf("   Common artists:  %d\n", commonArtists)
	fmt.Printf("   Similarity score: %.2f%%\n", score)
}
