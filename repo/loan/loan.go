package loan

import (
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

func (lr *LoanRepository) Loan(UserId int, BookId int, Address string) (_entities.Loan, int, error) {
	var loan _entities.Loan
	var user _entities.User
	var book _entities.Book
	tx := lr.database.Find(&book, BookId)
	if tx.Error != nil {
		return loan, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return loan, 0, tx.Error
	}
	err := lr.database.Find(&user, UserId)
	if err.Error != nil {
		return loan, 0, err.Error
	}
	if err.RowsAffected == 0 {
		return loan, 0, err.Error
	}
	if book.Qty > 0 {
		book.Qty = book.Qty - 1
		if book.Qty == 0 {
			book.Status = "book on loan"
		}
		lr.database.Save(&book)
	} else {
		return loan, 0, err.Error
	}
	loan.BookID = book.ID
	loan.UserID = user.ID
	loan.Address = Address
	if loan.Address == "" {
		return loan, 2, err.Error
	}
	loans := lr.database.Create(&loan)
	if loans.Error != nil {
		return loan, 0, loans.Error
	}
	if loans.RowsAffected == 0 {
		return loan, 0, loans.Error
	}
	return loan, int(loans.RowsAffected), nil
}

func (ur *LoanRepository) List() ([]_entities.Loan, error) {
	var list []_entities.Loan
	tx := ur.database.Find(&list)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return list, nil
}
