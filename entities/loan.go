package entities

import "gorm.io/gorm"

type Loan struct {
	gorm.Model
	UserID  uint
	BookID  uint
	Address string `json:"address" form:"address"`
}
