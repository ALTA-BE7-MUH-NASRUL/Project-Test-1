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
			Name    string `json:"name"`
			Book    string `json:"book"`
			Address string `json:"address"`
		}
		var returs returing
		c.Bind(&returs)
		retur, row, err := rh.returUseCase.Retur(returs.Name, returs.Book, returs.Address)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed "))
		}
		if row == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed returing book"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Success returing book", retur))
	}
}
