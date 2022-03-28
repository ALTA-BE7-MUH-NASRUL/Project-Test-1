package loan

import (
	_entities "latihan/coba-project/entities"
	_loanReapository "latihan/coba-project/repo/loan"
)

type LoanUseCase struct {
	loanRepository _loanReapository.LoanRepositoryInterface
}

func NewLoanUseCase(loanrepo _loanReapository.LoanRepositoryInterface) LoanUseCaseInterface {
	return &LoanUseCase{
		loanRepository: loanrepo,
	}
}

func (luc *LoanUseCase) Loan(Name string, Book string, Address string) (_entities.Loan, int, error) {
	loan, row, err := luc.loanRepository.Loan(Name, Book, Address)
	return loan, row, err
}