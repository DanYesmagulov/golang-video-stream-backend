package service

import (
	"github.com/DanYesmagulov/go-video-streaming/pkg/repository"
	"github.com/DanYesmagulov/go-video-streaming/pkg/store"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(category store.Category) (int, error) {
	return s.repo.Create(category)
}

func (s *CategoryService) GetAll() ([]store.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) GetById(id int) (store.Category, error) {
	return s.repo.GetById(id)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
