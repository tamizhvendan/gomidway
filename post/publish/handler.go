package publish

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tamizhvendan/gomidway/post"
	"github.com/tamizhvendan/gomidway/tag"
)

type Request struct {
	Title    string
	Body     string
	AuthorID uint
	Tags     []string
}

type Response struct {
	PostId uint
}

func NewPost(db *gorm.DB, req *Request) (*Response, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	newPost := &post.Post{
		AuthorID:    req.AuthorID,
		Title:       req.Title,
		Body:        req.Body,
		PublishedAt: time.Now().UTC(),
	}
	_, err := post.Create(tx, newPost)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for _, tagName := range req.Tags {
		t, err := tag.CreateIfNotExists(tx, tagName)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		err = post.AddTag(tx, newPost, t)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	res := tx.Commit()
	if res.Error != nil {
		return nil, res.Error
	}
	return &Response{PostId: newPost.ID}, nil
}
