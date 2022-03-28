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

func (luc *ReturUseCase) Retur(LoanId int, BookId int, Address string) (_entities.Retur, int, error) {
	retur, row, err := luc.returRepository.Retur(LoanId, BookId, Address)
	return retur, row, err
}
