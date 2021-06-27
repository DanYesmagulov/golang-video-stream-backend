package store

type Course struct {
	Id           int    `json:"id" db:"id"`
	Title        string `json:"title" db:"title"`
	Description  string `json:"description" db:"description"`
	Slug         string `json:"slug" db:"slug"`
	Language     string `json:"language" db:"language"`
	Likes        int    `json:"likes" db:"likes"`
	Dislikes     int    `json:"dislikes" db:"dislikes"`
	IsActive     bool   `json:"isActive" db:"is_active"`
	ImageUrl     string `json:"imageUrl" db:"image_url"`
	ArchiveUrl   string `json:"archiveUrl" db:"archive_url"`
	MaterialsUrl string `json:"materialsUrl" db:"materials_url"`
	CategoryId   int    `json:"categoryId" db:"category_id"`
	CreatedAt    string `json:"createdAt" db:"created_at"`
	UpdatedAt    string `json:"updatedAt" db:"updated_at"`
}

type UsersCourses struct {
	Id       int `db:"id"`
	UserId   int `db:"user_id"`
	CourseId int `db:"course_id"`
}

type Video struct {
	Id              int    `json:"id" db:"id"`
	Title           string `json:"title" db:"title"`
	Duration        string `json:"duration" db:"duration"`
	VideoOrder      int    `json:"videoOrder" db:"video_order"`
	VideoUrl        string `json:"videoUrl" db:"video_url"`
	FileName        string `json:"fileName" db:"file_name"`
	PreviewImageUrl string `json:"previewImageUrl" db:"preview_image_url"`
	CourseId        int    `json:"courseId" db:"course_id"`
}

type Comments struct {
	Id        int    `json:"id" db:"id"`
	Body      string `json:"body" db:"body"`
	CourseId  int    `json:"courseId" db:"course_id"`
	UserId    int    `json:"userId" db:"user_id"`
	CreatedAt string `json:"createdAt" db:"created_at"`
	UpdatedAt string `json:"updatedAt" db:"updated_at"`
}

type CourseRating struct {
	Id       int  `json:"id" db:"id"`
	UserId   int  `json:"userId"`
	CourseId int  `json:"courseId"`
	IsLike   bool `json:"isLike"`
}
