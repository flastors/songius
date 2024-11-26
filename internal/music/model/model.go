package model

type Music struct {
	ID          string `json:"id"`
	Song        string `json:"song"`
	Group       string `json:"artist"`
	ReleaseDate string `json:"release_date"`
	Link        string `json:"link"`
	Text        string `json:"text"`
}

func NewMusicModel(song, group, releaseDate, link, text string) *Music {
	return &Music{
		Song:        song,
		Group:       group,
		ReleaseDate: releaseDate,
		Link:        link,
		Text:        text,
	}
}
