# Spotify Match

`spotify-match` is a command-line application written in Go that compares two public Spotify playlists and evaluates how similar they are, based on shared tracks and artists. It uses the Spotify Web API with client credentials authentication and calculates a similarity score between the playlists.

## Features

- Prompts for two Spotify playlist links
- Automatically extracts playlist IDs from pasted URLs
- Authenticates using Spotify Web API (client credentials flow)
- Supports playlists of any size (pagination over 100 tracks)
- Compares:
  - Common tracks
  - Common artists
  - Similarity score in percent

## Example usage

```bash
go run main.go
```

```
Paste first playlist link:
> https://open.spotify.com/playlist/3wkJ5eRUxMezguGIFXjKjq

Paste second playlist link:
> https://open.spotify.com/playlist/6UJuK4qI1OpYUnF3tovYJ1

Analyzing playlists

Done!

Summary:
   Common songs:    43
   Common artists:  19
   Similarity score: 26.88%
```

## Installation

To install and run `spotify-match` locally, first clone the repository:

```bash
git clone https://github.com/robinho46/spotify-match.git
cd spotify-match
```

Then create a `.env` file in the root of the project with your Spotify API credentials:

```env
SPOTIFY_CLIENT_ID=your_client_id
SPOTIFY_CLIENT_SECRET=your_client_secret
```

You can obtain these credentials by registering an application at:  
https://developer.spotify.com/dashboard
> **Important:** Keep your `SPOTIFY_CLIENT_ID` and `SPOTIFY_CLIENT_SECRET` confidential.  
> Never commit them to version control or share them publicly. They grant access to your Spotify data.

Use the following as your redirect URI if needed:

```
http://127.0.0.1:8080/callback
```
