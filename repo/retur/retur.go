package retur

import (
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

func (rr *ReturRepository) Retur(LoanId int, BookId int, Address string) (_entities.Retur, int, error) {
	var retur _entities.Retur
	var user _entities.User
	var book _entities.Book
	var loan _entities.Loan
	err := rr.database.Find(&loan, LoanId)
	if err.Error != nil {
		return retur, 0, err.Error
	}
	if err.RowsAffected == 0 {
		return retur, 0, err.Error
	}
	if loan.Status == "returned" {
		return retur, 0, nil
	}
	tx := rr.database.Find(&book, BookId)
	if tx.Error != nil {
		return retur, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return retur, 0, tx.Error
	}
	loan.Status = "returned"
	rr.database.Save(&loan)
	rr.database.Find(&user, loan.UserID)
	book.Qty = book.Qty + 1
	book.Status = "book available"
	rr.database.Save(&book)
	retur.BookID = book.ID
	retur.UserID = user.ID
	retur.Address = Address
	returs := rr.database.Create(&retur)
	if returs.Error != nil {
		return retur, 0, returs.Error
	}
	if returs.RowsAffected == 0 {
		return retur, 0, returs.Error
	}
	return retur, int(returs.RowsAffected), nil
}
