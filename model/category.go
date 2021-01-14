package model

import (
	"vanwhebin/try-gin-vue/util"
)

type Category struct {
	ID uint `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"type:varchar(50);not null;unique"`
	Parent uint `json:"parent" gorm:"type:int(5)"`
	CreatedAt util.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt util.Time `json:"updated_at" gorm:"type:timestamp"`
}

