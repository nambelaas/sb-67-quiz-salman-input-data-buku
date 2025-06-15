package category

import (
	"database/sql"
	"errors"

	"github.com/sb-67-go-quiz-salman-input-data-buku/database/connection"
	"github.com/sb-67-go-quiz-salman-input-data-buku/structs"
)

func NewCategoryRepository() CategoryRepositoryInterface {
	return &CategoryRepository{}
}

func (c *CategoryRepository) CreateCategory(category structs.Category, responsibleUser string) error {
	query := `Insert into categories (name, created_by) values ($1,$2)`

	res, err := connection.DBConn.Exec(query, category.Name, responsibleUser)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal menambahkan data category")
	}

	return nil
}

func (c *CategoryRepository) GetAllCategory() (result []structs.Category, errorResult error) {
	query := `Select * from categories`

	rows, err := connection.DBConn.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var data = structs.Category{}
		var err = rows.Scan(&data.Id, &data.Name, &data.CreatedAt, &data.CreatedBy, &data.ModifiedAt, &data.ModifiedBy)
		if err != nil {
			return result, err
		}

		result = append(result, data)
	}

	return result, nil
}

func (c *CategoryRepository) GetCategoryById(id string) (result structs.Category, errorResult error) {
	query := `Select * from categories where id = $1`

	err := connection.DBConn.QueryRow(query, id).Scan(&result.Id, &result.Name, &result.CreatedAt, &result.CreatedBy, &result.ModifiedAt, &result.ModifiedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, errors.New("data category tidak ditemukan")
		}
		return result, err
	}

	return result, nil
}

func (c *CategoryRepository) GetBookInCategories(id string) (result []structs.Book, errorResult error) {
	query := `
		Select b.id, b.title, b.description, b.image_url, b.release_year, b.price, b.total_page, b.thickness, b.category_id, b.created_at, b.created_by, b.modified_at, b.modified_by 
		from books b
		join categories c on b.category_id = c.id
		where c.id = $1
	`

	rows, err := connection.DBConn.Query(query, id)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var book = structs.Book{}
		err := rows.Scan(
            &book.Id, &book.Title, &book.Description, &book.ImageUrl, &book.ReleaseYear, &book.Price,
            &book.TotalPage, &book.Thickness, &book.CategoryId, &book.CreatedAt, &book.CreatedBy,
            &book.ModifiedAt, &book.ModifiedBy,
        )
		if err != nil {
			return result, err
		}

		result = append(result, book)
	}

	return result, nil
}

func (c *CategoryRepository) DeleteCategory(id string) error {
	query := `Delete from categories where id = $1`

	res, err := connection.DBConn.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal menghapus category, id tidak ditemukan")
	}

	return nil
}
