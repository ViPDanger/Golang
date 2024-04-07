package postgres

import (
	"context"
	"log"
	"time"

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
		return 0, err
	}
	defer tx.Rollback(ctx)
	query := "insert into author (name) VALUES ($1)"
	_, err = tx.Exec(ctx, query, author.Name)

	if err != nil {
		return 0, err
	}
	query = "SELECT author.id FROM author WHERE author.name = $1"
	err = tx.QueryRow(context.Background(), query, author.Name).Scan(&i)
	if err != nil {
		return 0, err
	}
	tx.Commit(ctx)
	return i, err
}

func (r *Repository) Insert_Content(content structures.ContentDTO) (int, error) {
	var i int
	ctx := context.Background()
	tx, err := r.client.Begin(ctx)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer tx.Rollback(ctx)
	query := "insert into http_string (content,author_id,date) VALUES ($1,$2,$3)"
	_, err = tx.Exec(ctx, query, content.Content, content.Author_id, content.Date)

	if err != nil {
		log.Println(err)
		return 0, err
	}
	query = "SELECT http_string.id FROM http_string WHERE http_string.content = $1"
	err = tx.QueryRow(context.Background(), query, content.Content).Scan(&i)
	if err != nil {
		return 0, err
	}
	tx.Commit(ctx)
	return i, err
}

func (r *Repository) All_Authors() ([]structures.Author, error) {
	query := "SELECT * FROM author"
	authors := make([]structures.Author, 0)
	rows, err := r.client.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		authors = append(authors, structures.Author{})
		err = rows.Scan(&authors[i].ID, &authors[i].Name)
		if err != nil {
			return nil, err
		}
	}

	return authors, err
}

func (r *Repository) All_Content() ([]structures.Content, error) {
	query := "SELECT http_string.id,http_string.content,author.id,author.name,http_string.date FROM http_string inner join author on http_string.author_id = author.id"
	content := make([]structures.Content, 0)
	rows, err := r.client.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		content = append(content, structures.Content{})
		var time time.Time
		err = rows.Scan(&content[i].ID, &content[i].Content, &content[i].Author.ID, &content[i].Author.Name, &time)
		content[i].Date = time.Format("02.01.2006  15:04:05")
	}
	return content, err
}

func (r *Repository) Find_Author(id string) (structures.Author, error) {
	var author structures.Author
	query := "SELECT * FROM author WHERE id = $1"
	err := r.client.QueryRow(context.Background(), query, id).Scan(&author.ID, &author.Name)

	return author, err
}

func (r *Repository) Find_Content(id string) (structures.Content, error) {
	var content structures.Content
	query := "SELECT http_string.id,http_string.content,author.id,author.name,http_string.date FROM http_string inner join author on http_string.author_id = author.id  WHERE http_string.id = $1"
	var time time.Time
	err := r.client.QueryRow(context.Background(), query, id).Scan(&content.ID, &content.Content, &content.Author.ID, &content.Author.Name, &time)
	content.Date = time.Format("02.01.2006  15:04:05")
	return content, err
}
func (r *Repository) Delete_Author(id string) error {
	ctx := context.Background()
	tx, err := r.client.Begin(ctx)
	if err != nil {
		config.Err_log(err)
		return err
	}
	defer tx.Rollback(ctx)
	query := "DELETE FROM author WHERE author.id = $1"
	_, err = tx.Exec(ctx, query, id)

	if err != nil {
		return err
	}
	tx.Commit(ctx)
	return err
}

func (r *Repository) Delete_Content(id string) error {
	ctx := context.Background()
	tx, err := r.client.Begin(ctx)
	if err != nil {
		config.Err_log(err)
		return err
	}
	defer tx.Rollback(ctx)
	query := "DELETE FROM http_string WHERE http_string.id = $1"
	_, err = tx.Exec(ctx, query, id)

	if err != nil {
		return err
	}
	tx.Commit(ctx)
	return err
}
