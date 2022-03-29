package retur

import (
	"latihan/coba-project/delivery/helper"
	_returUseCase "latihan/coba-project/usecase/retur"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReturHandler struct {
	returUseCase _returUseCase.ReturUseCaseInterface
}

func NewReturHandler(returUseCase _returUseCase.ReturUseCaseInterface) *ReturHandler {
	return &ReturHandler{
		returUseCase: returUseCase,
	}
}

func (rh *ReturHandler) ReturingHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		type returing struct {
			LoanId  int    `json:"LoanId"`
			BookId  int    `json:"BookId"`
			Address string `json:"address"`
		}
		var returs returing
		c.Bind(&returs)
		retur, row, err := rh.returUseCase.Retur(returs.LoanId, returs.BookId, returs.Address)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to returning book"))
		}
		if row == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("the book has been returned"))
		}
		if row == 2 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("wrong LoanId"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Success returning book", retur))
	}
}
