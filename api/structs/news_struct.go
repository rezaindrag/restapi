package structs

// News struct
type News struct {
	ID          int    `json:"id" validate:""`
	Title       string `json:"title" validate:"required,min=5,max=50"`
	Description string `json:"description" validate:"required,min=10,max=100"`
	Thumbnail   string `json:"thumbnail" validate:"required,url"`
	Author      string `json:"author" validate:"required"`
	PublishDate string `json:"publish_date" validate:""`
}
