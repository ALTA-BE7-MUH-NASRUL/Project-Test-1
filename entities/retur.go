package entities

import "gorm.io/gorm"

type Retur struct {
	gorm.Model
	UserID  uint
	BookID  uint
	Address string `json:"address" form:"address"`
}
