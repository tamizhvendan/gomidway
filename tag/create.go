package tag

import "github.com/jinzhu/gorm"

func CreateIfNotExists(db *gorm.DB, tagName string) (*Tag, error) {
	var tag Tag
	res := db.FirstOrCreate(&tag, Tag{Name: tagName})
	if res.Error != nil {
		return nil, res.Error
	}
	return &tag, nil
}
