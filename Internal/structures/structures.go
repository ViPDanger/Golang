package structures

type Get_Page struct {
	Title            string
	SearchAllAuthor  string
	SearchAllContent string
	SearchAuthorID   string
	SearchContentID  string
	InputAuthor      string
	InputContent     string
	DeleteAuthorID   string
	DeleteContentID  string
}

type Result_Page struct {
	Title      string
	First_Line []string
	Data       any
	BackButton string
}

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Content struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  Author `json:"author"`
	Date    string `json:"date"`
}

type ContentDTO struct {
	Content   string `json:"content"`
	Author_id int    `json:"author_id"`
	Date      string `json:"date"`
}

type AuthorDTO struct {
	Name string `json:"name"`
}
