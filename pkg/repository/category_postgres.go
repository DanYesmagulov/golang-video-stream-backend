package repository

import (
	"fmt"
	"github.com/DanYesmagulov/go-video-streaming/pkg/store"
	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) Create(category store.Category) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createCategoryQuery := fmt.Sprintf("INSERT INTO %s (title, description, slug, image_url, parent_id) VALUES ($1, $2, $3, $4, $5) RETURNING id", categoryTable)
	row := tx.QueryRow(createCategoryQuery, category.Title, category.Description, category.Slug, category.ImageUrl, category.ParentId)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *CategoryPostgres) GetAll() ([]store.Category, error) {
	var categories []store.Category
	query := fmt.Sprintf("SELECT * FROM %s", categoryTable)
	err := r.db.Select(&categories, query)

	return categories, err
}

func (r *CategoryPostgres) GetById(id int) (store.Category, error) {
	var category store.Category
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", categoryTable)
	err := r.db.Get(&category, query, id)

	return category, err
}

func (r *CategoryPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", categoryTable)
	_, err := r.db.Exec(query, id)

	return err
}
