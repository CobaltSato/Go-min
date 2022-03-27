package post

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Name    string
	Message string
}
