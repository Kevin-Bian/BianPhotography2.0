package models

type Photo struct {
	ID          int64  `json:"photoid"`
	CollageID   string `json:"collageid"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	Description string `json:"description"`
}
