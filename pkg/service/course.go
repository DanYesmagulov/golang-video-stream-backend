package service

import (
	"github.com/DanYesmagulov/go-video-streaming/pkg/repository"
	"github.com/DanYesmagulov/go-video-streaming/pkg/store"
)

type CourseService struct {
	repo repository.Course
}

func NewCourseService(repo repository.Course) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) Create(userId int, category store.Course) (int, error) {
	return s.repo.Create(userId, category)
}

func (s *CourseService) GetAll() ([]store.Course, error) {
	return s.repo.GetAll()
}

func (s *CourseService) GetById(id int) (store.Course, error) {
	return s.repo.GetById(id)
}

func (s *CourseService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}
