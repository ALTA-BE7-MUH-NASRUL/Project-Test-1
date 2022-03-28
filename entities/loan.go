package entities

import "gorm.io/gorm"

type Loan struct {
	gorm.Model
	UserID  uint
	BookID  uint
	Status  string `json:"status" form:"status"`
	Address string `json:"address" form:"address"`
}
