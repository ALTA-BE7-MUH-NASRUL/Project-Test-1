package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `json:"name" form:"name"`
	Email    string  `json:"email" form:"email"`
	Password string  `json:"password" form:"password"`
	Address  string  `json:"address" form:"address"`
	Retur    []Retur `gorm:"foreignKey:UserID;references:ID"`
	Loan     []Loan  `gorm:"foreignKey:UserID;references:ID"`
	Book     []Book  `gorm:"foreignKey:UserID;references:ID"`
}
