package loan

import _entities "latihan/coba-project/entities"

type LoanRepositoryInterface interface {
	Loan(UserId int, BookId int, Address string) (_entities.Loan, int, error)
}
