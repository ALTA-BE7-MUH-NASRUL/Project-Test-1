package retur

import (
	"errors"
	_entities "latihan/coba-project/entities"

	"gorm.io/gorm"
)

type ReturRepository struct {
	database *gorm.DB
}

func NewReturRepository(db *gorm.DB) *ReturRepository {
	return &ReturRepository{
		database: db,
	}
}

func (rr *ReturRepository) Retur(Name string, Book string, Address string) (_entities.Retur, int, error) {
	var retur _entities.Retur
	var user _entities.User
	var book _entities.Book
	tx := rr.database.Where("title = ?", Book).Find(&book)
	if tx.Error != nil {
		return retur, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return retur, 0, errors.New("book not found")
	}

	err := rr.database.Where("name=?", Name).Find(&user)
	if err.Error != nil {
		return retur, 0, err.Error
	}
	if err.RowsAffected == 0 {
		return retur, 0, errors.New("user not found")
	}
	book.Qty = book.Qty + 1
	rr.database.Save(&retur)
	retur.BookID = book.ID
	retur.UserID = user.ID
	retur.Address = Address
	returs := rr.database.Create(&retur)
	if returs.Error != nil {
		return retur, 0, returs.Error
	}
	if returs.RowsAffected == 0 {
		return retur, 0, errors.New("something wrong")
	}
	return retur, int(returs.RowsAffected), nil
}
