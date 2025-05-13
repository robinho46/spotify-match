package models

type Track struct {
	Name   string
	Artist string
	ID     string
}

type PlaylistResponse struct {
	Items []struct {
		Track struct {
			Name    string `json:"name"`
			ID      string `json:"id"`
			Artists []struct {
				Name string `json:"name"`
			} `json:"artists"`
		} `json:"track"`
	} `json:"items"`
	Total int `json:"total"`
}
