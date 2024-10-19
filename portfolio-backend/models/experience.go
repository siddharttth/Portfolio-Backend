package models

type Experience struct {
	Title        string   `json:"title"`
	Company      string   `json:"company"`
	Location     string   `json:"location"`
	Period       string   `json:"period"`
	Achievements []string `json:"achievements"`
}
