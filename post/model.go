package post

type Post struct {
	ID    uint `gorm:"primary_key"`
	Title string
	Body  string
}
