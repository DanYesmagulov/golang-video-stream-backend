package repository

import (
	"fmt"
	"github.com/DanYesmagulov/go-video-streaming/pkg/store"
	"github.com/jmoiron/sqlx"
)

type CoursePostgres struct {
	db *sqlx.DB
}

func NewCoursePostgres(db *sqlx.DB) *CoursePostgres {
	return &CoursePostgres{db: db}
}

func (r *CoursePostgres) Create(userId int, course store.Course) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createCourseQuery := fmt.Sprintf("INSERT INTO %s (title, description, slug, language, image_url, category_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", courseTable)
	row := tx.QueryRow(createCourseQuery, course.Title, course.Description, course.Slug, course.Language, course.ImageUrl, course.CategoryId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersCoursesQuery := fmt.Sprintf("INSERT INTO %s (user_id, course_id) VALUES ($1, $2)", usersCoursesTable)
	_, err = tx.Exec(createUsersCoursesQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *CoursePostgres) GetAll() ([]store.Course, error) {
	var course []store.Course
	query := fmt.Sprintf("SELECT * FROM %s", courseTable)
	err := r.db.Select(&course, query)

	return course, err
}

func (r *CoursePostgres) GetById(id int) (store.Course, error) {
	var course store.Course
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", courseTable)
	err := r.db.Get(&course, query, id)

	return course, err
}

func (r *CoursePostgres) Delete(userId, id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	deleteUsersCoursesQuery := fmt.Sprintf("DELETE FROM %s WHERE course_id = $1 AND user_id = $2", usersCoursesTable)
	_, err = tx.Exec(deleteUsersCoursesQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	deleteCourseQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", courseTable)
	_, err = tx.Exec(deleteCourseQuery, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
