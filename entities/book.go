package entities

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	UserID        uint
	Title         string  `json:"title" form:"title"`
	Author        string  `json:"author" form:"author"`
	Publisher     string  `json:"publisher" form:"publisher"`
	Owner_Name    string  `json:"owner_name" form:"owner_name"`
	Owner_Address string  `json:"owner_address" form:"owner_address"`
	Qty           int     `json:"qty" form:"qty"`
	Status        string  `json:"status" form:"status"`
	Retur         []Retur `gorm:"foreignKey:BookID;references:ID"`
	Loan          []Loan  `gorm:"foreignKey:BookID;references:ID"`
}
