package system

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Id        *uint          `gorm:"primarykey" json:"id"`
	CreatedAt *time.Time     `json:"createdAt"`
	UpdatedAt *time.Time     `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func setDefault[T any](ptr **T, defaultValue T) {
	if *ptr == nil {
		*ptr = &defaultValue
	}
}
