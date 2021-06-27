package service

import (
	"context"
	"github.com/DanYesmagulov/go-video-streaming/pkg/repository"
	"github.com/DanYesmagulov/go-video-streaming/pkg/storage"
	"github.com/DanYesmagulov/go-video-streaming/pkg/store"
	"io"
)

type Authorization interface {
	CreateUser(user store.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
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
	Upload(ctx context.Context, inp UploadInput) (string, error)
}

type Service struct {
	Authorization
	Category
	Course
	File
}

type UploadInput struct {
	File          io.Reader
	FileName      string
	FileExtension string
	Size          int64
	ContentType   string
	Type          FileType
	CourseId int
}

type Deps struct {
	Repos           *repository.Repository
	StorageProvider storage.Provider
}

func NewService(deps Deps) *Service {
	return &Service{
		Authorization: NewAuthService(deps.Repos.Authorization),
		Category:      NewCategoryService(deps.Repos.Category),
		Course:        NewCourseService(deps.Repos.Course),
		File:          NewFilesService(deps.StorageProvider, deps.Repos.File),
	}
}
