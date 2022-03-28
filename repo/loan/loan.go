package loan

import (
	"errors"
	_entities "latihan/coba-project/entities"

	"gorm.io/gorm"
)

type LoanRepository struct {
	database *gorm.DB
}

func NewLoanRepository(db *gorm.DB) *LoanRepository {
	return &LoanRepository{
		database: db,
	}
}

func (lr *LoanRepository) Loan(Name string, Book string, Address string) (_entities.Loan, int, error) {
	var loan _entities.Loan
	var user _entities.User
	var book _entities.Book
	tx := lr.database.Where("title = ?", Book).Find(&book)
	if tx.Error != nil {
		return loan, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return loan, 0, errors.New("book not found")
	}

	err := lr.database.Where("name=?", Name).Find(&user)
	if err.Error != nil {
		return loan, 0, err.Error
	}
	if err.RowsAffected == 0 {
		return loan, 0, errors.New("user not found")
	}
	if book.Qty > 0 {
		book.Qty = book.Qty - 1
		lr.database.Save(&book)
	} else {
		return loan, 0, errors.New("book on loan")
	}
	loan.BookID = book.ID
	loan.UserID = user.ID
	loan.Address = Address
	loans := lr.database.Create(&loan)
	if loans.Error != nil {
		return loan, 0, loans.Error
	}
	if loans.RowsAffected == 0 {
		return loan, 0, errors.New("something wrong")
	}
	return loan, int(loans.RowsAffected), nil
}
