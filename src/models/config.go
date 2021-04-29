package models

type Config struct {
	Code        string  `json:"code"`
	Description *string `json:"description"`
	Path        string  `json:"path"`
}
