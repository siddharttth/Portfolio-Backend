package models

type Project struct {
	Name      string   `json:"name"`
	Period    string   `json:"period"`
	TechStack string   `json:"techStack"`
	Points    []string `json:"points"`
}
