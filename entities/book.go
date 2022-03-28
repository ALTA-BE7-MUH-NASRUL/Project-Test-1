package entities

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title         string  `json:"title" form:"title"`
	Author        string  `json:"author" form:"author"`
	Publisher     string  `json:"publisher" form:"publisher"`
	Owner_Name    string  `json:"owner_name" form:"owner_name"`
	Owner_Address string  `json:"owner_address" form:"owner_address"`
	Qty           int     `json:"qty" form:"qty"`
	Retur         []Retur `gorm:"foreignKey:BookID;references:ID"`
	Loan          []Loan  `gorm:"foreignKey:BookID;references:ID"`
}
