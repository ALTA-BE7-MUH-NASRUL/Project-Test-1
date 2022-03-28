package loan

import _entities "latihan/coba-project/entities"

type LoanUseCaseInterface interface {
	Loan(Name string, Book string, Address string) (_entities.Loan, int, error)
}
