package retur

import _entities "latihan/coba-project/entities"

type ReturRepositoryInterface interface {
	Retur(Name string, Book string, Address string) (_entities.Retur, int, error)
}
