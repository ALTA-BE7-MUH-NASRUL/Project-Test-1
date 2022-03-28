package loan

import (
	"latihan/coba-project/delivery/helper"
	_loanUseCase "latihan/coba-project/usecase/loan"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoanHandler struct {
	loanUseCase _loanUseCase.LoanUseCaseInterface
}

func NewLoanHandler(loanUseCase _loanUseCase.LoanUseCaseInterface) *LoanHandler {
	return &LoanHandler{
		loanUseCase: loanUseCase,
	}
}

func (lh *LoanHandler) LoaningHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		type loaning struct {
			Name    string `json:"name"`
			Book    string `json:"book"`
			Address string `json:"address"`
		}
		var loans loaning
		c.Bind(&loans)
		loan, row, err := lh.loanUseCase.Loan(loans.Name, loans.Book, loans.Address)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed "))
		}
		if row == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("book on loan"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Success loaning book", loan))
	}
}
