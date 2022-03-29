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
			UserId  int    `json:"UserId"`
			BookId  int    `json:"BookId"`
			Address string `json:"address"`
		}
		var loans loaning
		c.Bind(&loans)
		loan, row, err := lh.loanUseCase.Loan(loans.UserId, loans.BookId, loans.Address)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to loaning book"))
		}
		if row == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("book on loan"))
		}
		if row == 2 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("please insert your address"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Success loaning book", loan))
	}
}

func (uh *LoanHandler) ListHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		list, err := uh.loanUseCase.List()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to fetch data"))
		}
		if len(list) == 0 {
			return c.JSON(http.StatusOK, helper.ResponseSuccess("Data not exist", list))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all data", list))
	}
}
