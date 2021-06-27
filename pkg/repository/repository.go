package repository

import (
	"github.com/DanYesmagulov/go-video-streaming/pkg/store"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user store.User) (int, error)
	GetUser(email, password string) (store.User, error)
}

type Category interface {
	Create(category store.Category) (int, error)
	GetAll() ([]store.Category, error)
	GetById(id int) (store.Category, error)
	Delete(id int) error
}

type Course interface {
	Create(userId int, course store.Course) (int, error)
	GetAll() ([]store.Course, error)
	GetById(id int) (store.Course, error)
	Delete(userId, id int) error
}

type File interface {
	Create(video store.Video) (store.Video, error)
}

type Repository struct {
	Authorization
	Category
	Course
	File
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Category:      NewCategoryPostgres(db),
		Course:        NewCoursePostgres(db),
		File: NewFilePostgres(db),
	}
}
