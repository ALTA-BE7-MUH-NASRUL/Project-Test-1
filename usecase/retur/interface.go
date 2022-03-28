package retur

import _entities "latihan/coba-project/entities"

type ReturUseCaseInterface interface {
	Retur(LoanId int, BookId int, Address string) (_entities.Retur, int, error)
}
