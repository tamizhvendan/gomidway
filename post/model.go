package post

import (
	"time"

	"github.com/tamizhvendan/gomidway/tag"
)

const (
	UniqueConstraintTitle = "posts_title_key"
	AssociationTags       = "Tags"
)

type Post struct {
	ID          uint
	Title       string
	Body        string
	AuthorID    uint
	Tags        []tag.Tag `gorm:"many2many:posts_tags;"`
	PublishedAt time.Time
}

type TitleDuplicateError struct{}

func (e *TitleDuplicateError) Error() string {
	return "title already exists"
}
