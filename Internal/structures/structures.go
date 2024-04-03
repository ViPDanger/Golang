package structures

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Content struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Author  Author `json:"author"`
	Date    string `json:"date"`
}

type ContentDTO struct {
	Content   string `json:"content"`
	Author_id string `json:"author_id"`
	Date      string `json:"date"`
}

type AuthorDTO struct {
	Name string `json:"name"`
}
