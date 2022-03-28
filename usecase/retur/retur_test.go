package retur

import (
	"fmt"
	_entities "latihan/coba-project/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetur(t *testing.T) {
	t.Run("TestLoanSuccess", func(t *testing.T) {
		returUseCase := NewReturUseCase(mockLoanRepository{})
		data, row, err := returUseCase.Retur(1, 1, "Address")
		assert.Nil(t, err)
		assert.Equal(t, "success", data.Address)
		assert.Equal(t, 1, row)
	})

	t.Run("TestLoanError", func(t *testing.T) {
		returUseCase := NewReturUseCase(mockLoanRepositoryError{})
		data, row, err := returUseCase.Retur(1, 0, "Address")
		assert.NotNil(t, err)
		assert.Nil(t, nil, data.Address)
		assert.Equal(t, 0, row)

	})
}

type mockLoanRepository struct{}

func (m mockLoanRepository) Retur(UserId int, BookId int, Address string) (_entities.Retur, int, error) {
	return _entities.Retur{Address: "success"}, 1, nil
}

type mockLoanRepositoryError struct{}

func (m mockLoanRepositoryError) Retur(UserId int, BookId int, Address string) (_entities.Retur, int, error) {
	return _entities.Retur{}, 0, fmt.Errorf("error")
}
