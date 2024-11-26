package dto

type GetInfoReq struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}
type GetInfoRes struct {
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
