package retur

import _entities "latihan/coba-project/entities"

type ReturRepositoryInterface interface {
	Retur(LoanId int, BookId int, Address string) (_entities.Retur, int, error)
}
