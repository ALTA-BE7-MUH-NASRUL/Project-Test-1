package loan

import (
	"fmt"
	_entities "latihan/coba-project/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoan(t *testing.T) {
	t.Run("TestLoanSuccess", func(t *testing.T) {
		loanUseCase := NewLoanUseCase(mockLoanRepository{})
		data, row, err := loanUseCase.Loan(1, 1, "Address")
		assert.Nil(t, err)
		assert.Equal(t, "success", data.Address)
		assert.Equal(t, 1, row)
	})

	t.Run("TestLoanError", func(t *testing.T) {
		loanUseCase := NewLoanUseCase(mockLoanRepositoryError{})
		data, row, err := loanUseCase.Loan(1, 0, "Address")
		assert.NotNil(t, err)
		assert.Nil(t, nil, data.Address)
		assert.Equal(t, 0, row)

	})
}
func TestList(t *testing.T) {
	t.Run("TestListSuccess", func(t *testing.T) {
		listUseCase := NewLoanUseCase(mockLoanRepository{})
		data, err := listUseCase.List()
		assert.Nil(t, err)
		assert.Equal(t, "success", data[0].Address)
	})

	t.Run("TestListError", func(t *testing.T) {
		listUseCase := NewLoanUseCase(mockLoanRepositoryError{})
		data, err := listUseCase.List()
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
	})
}

type mockLoanRepository struct{}

func (m mockLoanRepository) Loan(UserId int, BookId int, Address string) (_entities.Loan, int, error) {
	return _entities.Loan{Address: "success"}, 1, nil
}
func (m mockLoanRepository) List() ([]_entities.Loan, error) {
	return []_entities.Loan{{Address: "success"}}, nil
}

type mockLoanRepositoryError struct{}

func (m mockLoanRepositoryError) Loan(UserId int, BookId int, Address string) (_entities.Loan, int, error) {
	return _entities.Loan{}, 0, fmt.Errorf("error")
}
func (m mockLoanRepositoryError) List() ([]_entities.Loan, error) {
	return nil, fmt.Errorf("error")
}
