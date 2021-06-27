package repository

import (
	"fmt"
	"github.com/DanYesmagulov/go-video-streaming/pkg/store"
	"github.com/jmoiron/sqlx"
)

type FilePostgres struct {
	db *sqlx.DB
}

func NewFilePostgres(db *sqlx.DB) *FilePostgres {
	return &FilePostgres{db: db}
}

func (r *FilePostgres) Create(video store.Video) (store.Video, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return store.Video{}, err
	}

	var file store.Video
	createCourseQuery := fmt.Sprintf("INSERT INTO %s (video_url, file_name, preview_image_url, course_id) VALUES ($1, $2, $3, $4) RETURNING id", videoTable)
	_, err = tx.Exec(createCourseQuery, video.VideoUrl, video.FileName, video.PreviewImageUrl, video.CourseId)
	if err != nil {
		tx.Rollback()
		return store.Video{}, err
	}

	return file, tx.Commit()
}