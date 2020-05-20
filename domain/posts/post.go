package posts

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title string  `gorm:"type:varchar(100)" json:"title" binding:"required"`
	Des   string  `gorm:"type:varchar(500)" json:"des"   binding:"required"`

}