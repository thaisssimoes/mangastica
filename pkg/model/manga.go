package model

type Manga struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	CoverLink   string `json:"coverLink"`
}
