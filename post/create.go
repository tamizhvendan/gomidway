package post

import (
	"github.com/jinzhu/gorm"
	"github.com/tamizhvendan/gomidway/postgres"
)

func Create(db *gorm.DB, post *Post) (uint, error) {
	res := db.Create(post)
	if res.Error != nil {
		if postgres.IsUniqueConstraintError(res.Error, UniqueConstraintTitle) {
			return 0, &TitleDuplicateError{}
		}
		return 0, res.Error
	}
	return post.ID, nil
}
