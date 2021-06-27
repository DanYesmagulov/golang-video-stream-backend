package service

import (
	"context"
	"fmt"
	"github.com/DanYesmagulov/go-video-streaming/pkg/repository"
	"github.com/DanYesmagulov/go-video-streaming/pkg/storage"
	"github.com/DanYesmagulov/go-video-streaming/pkg/store"
	"github.com/google/uuid"
	"time"
)

type FileType int

const (
	FileTypeImage FileType = iota
	FileTypeVideo
)

var folders = map[FileType]string{
	FileTypeImage: "category-images",
	FileTypeVideo: "video-file",
}

const (
	uploadTimeout = time.Minute
)

type FilesService struct {
	repo repository.File
	storage storage.Provider
}

func NewFilesService(storage storage.Provider, repo repository.File) *FilesService {
	return &FilesService{storage: storage, repo: repo}
}

func (s *FilesService) Upload(ctx context.Context, inp UploadInput) (string, error) {
	ctx, clFn := context.WithTimeout(ctx, uploadTimeout)
	defer clFn()

	fmt.Println(s.generateFilename(inp))

	url, _ := s.storage.Upload(ctx, storage.UploadInput{
		File:        inp.File,
		Size:        inp.Size,
		ContentType: inp.ContentType,
		Name:        s.generateFilename(inp),
	})


	s.repo.Create(store.Video{
		VideoUrl:        url,
		FileName:        s.generateFilename(inp),
		PreviewImageUrl: "",
		CourseId:        inp.CourseId,
	})

	return url, nil
}

func (s *FilesService) generateFilename(inp UploadInput) string {
	filename := fmt.Sprintf("%s_%s.%s", inp.FileName, uuid.New().String(), inp.FileExtension)

	folder := folders[inp.Type]

	return fmt.Sprintf("%s/%s", folder, filename)
}
