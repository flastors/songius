package model

type CreateMusicDTO struct {
	Song  string `json:"song"`
	Group string `json:"group"`
}

type UpdateMusicDTO struct {
	Song        string `json:"song"`
	Group       string `json:"artist"`
	ReleaseDate string `json:"release_date"`
	Link        string `json:"link"`
	Text        string `json:"text"`
}

type APIMusicDTO struct {
	Song        string `json:"song"`
	Group       string `json:"group"`
	ReleaseDate string `json:"release_date"`
	Link        string `json:"link"`
	Text        string `json:"text"`
}
