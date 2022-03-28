package loan

import _entities "latihan/coba-project/entities"

type LoanUseCaseInterface interface {
	Loan(UserId int, BookId int, Address string) (_entities.Loan, int, error)
	List() ([]_entities.Loan, error)
}
