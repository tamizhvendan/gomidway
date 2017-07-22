package comment

type Comment struct {
	ID   uint `gorm:"primary_key"`
	Body string
}
