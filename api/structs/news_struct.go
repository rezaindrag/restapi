package structs

// News struct
type News struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
}
