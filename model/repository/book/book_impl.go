package book

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/sb-67-go-quiz-salman-input-data-buku/database/connection"
	"github.com/sb-67-go-quiz-salman-input-data-buku/structs"
)

func NewBookRepository() BookRepositoryInterface {
	return &BookRepository{}
}

func (c *BookRepository) CreateBook(book structs.Book, responsibleUser string) error {
	var thickness string
	if book.TotalPage > 100 {
		thickness = "tebal"
	} else {
		thickness = "tipis"
	}

	query := `Insert into books (title, category_id, description, image_url, release_year,price,total_page,thickness, created_by) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	res, err := connection.DBConn.Exec(query, book.Title, book.CategoryId, book.Description, book.ImageUrl, book.ReleaseYear, book.Price, book.TotalPage, thickness, responsibleUser)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return errors.New("kategori tidak ditemukan, pastikan category_id valid")
		}
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal menambahkan data book")
	}

	return nil
}

func (c *BookRepository) GetAllBook() (result []structs.Book, errorResult error) {
	query := `Select * from books`

	rows, err := connection.DBConn.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var data = structs.Book{}
		var err = rows.Scan(&data.Id, &data.Title, &data.Description, &data.ImageUrl, &data.ReleaseYear, &data.Price, &data.TotalPage, &data.Thickness, &data.CategoryId, &data.CreatedAt, &data.CreatedBy, &data.ModifiedAt, &data.ModifiedBy)
		if err != nil {
			return result, err
		}

		result = append(result, data)
	}

	return result, nil
}

func (c *BookRepository) GetBookById(id string) (result structs.Book, errorResult error) {
	query := `Select * from books where id = $1`

	err := connection.DBConn.QueryRow(query, id).Scan(&result.Id, &result.Title, &result.Description, &result.ImageUrl, &result.ReleaseYear, &result.Price, &result.TotalPage, &result.Thickness, &result.CategoryId, &result.CreatedAt, &result.CreatedBy, &result.ModifiedAt, &result.ModifiedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, errors.New("data book tidak ditemukan")
		}
		return result, err
	}

	return result, nil
}

func (c *BookRepository) DeleteBook(id string) error {
	query := `Delete from books where id = $1`

	res, err := connection.DBConn.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal menghapus book, id tidak ditemukan")
	}

	return nil
}
