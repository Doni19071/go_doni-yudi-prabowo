package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name       string `json:"name" form:"name"`
	Pengarang  string `json:"pengarang" form:"pengarang"`
	TahunRilis string `json:"tahunrilis" form:"tahunrilis"`
}
