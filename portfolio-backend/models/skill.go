package models

type Skill struct {
	Category string   `json:"category"`
	Items    []string `json:"items"`
}
