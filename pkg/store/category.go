package store

type Category struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Slug        string `json:"slug" db:"slug"`
	ImageUrl    string `json:"imageUrl" db:"image_url"`
	ParentId    int    `json:"parentId" db:"parent_id"`
}

type UpdateCategory struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Slug        *string `json:"slug"`
	ImageUrl    *string `json:"imageUrl"`
	ParentId    *int    `json:"parentId"`
}
