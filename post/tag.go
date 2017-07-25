package post

import (
	"github.com/jinzhu/gorm"
	"github.com/tamizhvendan/gomidway/tag"
)

func AddTag(db *gorm.DB, post *Post, tag *tag.Tag) error {
	res := db.Model(post).Association(AssociationTags).Append(tag)
	return res.Error
}
