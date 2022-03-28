package retur

import (
	_entities "latihan/coba-project/entities"
	_returReapository "latihan/coba-project/repo/retur"
)

type ReturUseCase struct {
	returRepository _returReapository.ReturRepositoryInterface
}

func NewReturUseCase(returrepo _returReapository.ReturRepositoryInterface) ReturUseCaseInterface {
	return &ReturUseCase{
		returRepository: returrepo,
	}
}

func (luc *ReturUseCase) Retur(Name string, Book string, Address string) (_entities.Retur, int, error) {
	retur, row, err := luc.returRepository.Retur(Name, Book, Address)
	return retur, row, err
}
