package models

type Education struct {
	Degree    string `json:"degree"`
	Institute string `json:"institute"`
	CGPA      string `json:"cgpa"`
	Year      string `json:"year"`
}
