package postgress

import (
	"context"

	"github.com/ViPDanger/Golang/Internal/config"
	"github.com/ViPDanger/Golang/Internal/structures"
)

// Описание репозитория клиента
type Repository struct {
	client Client
}

func NewRepository(client Client) *Repository {
	return &Repository{
		client: client,
	}
}

// Функции репозитория
func (r *Repository) Insert_Author(author structures.AuthorDTO) (int, error) {
	var i int
	ctx := context.Background()
	tx, err := r.client.Begin(ctx)
	if err != nil {
		config.Err_log(err)
		return 0, err
	}
	defer tx.Rollback(ctx)
	query := "insert into author (name) VALUES ($1)"
	_, err = tx.Exec(ctx, query, author.Name)

	if err != nil {
		config.Err_log(err)
		return 0, err
	}
	query = "SELECT author.id FROM author WHERE author.name = $1"
	err = r.client.QueryRow(context.Background(), query, author.Name).Scan(&i)
	config.Err_log(err)
	return i, err
}

func (r *Repository) Insert_Content(content structures.ContentDTO) int {
	var i int
	ctx := context.Background()
	tx, err := r.client.Begin(ctx)
	if err != nil {
		config.Err_log(err)
		return 0, err
	}
	defer tx.Rollback(ctx)
	query := "insert into http_string (content,author_id,date) VALUES ($1,$2,$3)"
	_, err = tx.Exec(ctx, query, content.Content, content.Author_id, content.Date)

}

func (r *Repository) All_Authors() (*[]structures.Author, error)

func (r *Repository) All_Content() (*[]structures.Content, error)

func (r *Repository) Find_Author(id int) (*structures.Author, error) {
	var author structures.Author
	query := "SELECT * FROM author WHERE author.id = $1"
	err := r.client.QueryRow(context.Background(), query, id).Scan(&author.ID, &author.Name)
	config.Err_log(err)
	return &author, err
}

func (r *Repository) Find_Content(id int) (*structures.Content, error) {
	var content structures.Content
	query := "SELECT http_string.id,http_string.content,author.id,author.name FROM http_string inner join author on http_string.author_id = author_id  WHERE http_string.id = $1"
	err := r.client.QueryRow(context.Background(), query, id).Scan(&content.ID, &content.Content, &content.Author.ID, &content.Author.Name)
	config.Err_log(err)
	return &content, err
}
