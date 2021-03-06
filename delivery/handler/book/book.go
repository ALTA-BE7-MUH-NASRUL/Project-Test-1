package book

import (
	"latihan/coba-project/delivery/helper"
	_middlewares "latihan/coba-project/delivery/middleware"
	_entities "latihan/coba-project/entities"
	_bookUseCase "latihan/coba-project/usecase/book"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookUseCase _bookUseCase.BookUseCaseInterface
}

func NewBookHandler(bookUseCase _bookUseCase.BookUseCaseInterface) *BookHandler {
	return &BookHandler{
		bookUseCase: bookUseCase,
	}
}

func (bh *BookHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		books, err := bh.bookUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed fetch data"))
		}
		if len(books) == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Data not exist"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all data", books))
	}
}

func (bh *BookHandler) GetBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		books, rows, err := bh.bookUseCase.GetBook(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not exist"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("succes get data", books))
	}
}

func (bh *BookHandler) DeleteBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var loan _entities.Loan
		var user _entities.Book
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		_, rows, err := bh.bookUseCase.GetBook(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not exist"))
		}
		if idToken != int(user.UserID) {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("cannot delete book id"))
		}
		if loan.BookID == uint(id) && loan.Status == "book on loan" {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("book has not been returned"))
		}
		_, err = bh.bookUseCase.DeleteBook(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed delete data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success delete data"))
	}
}
func (bh *BookHandler) CreateBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var book _entities.Book
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		book.UserID = uint(idToken)
		c.Bind(&book)
		books, err := bh.bookUseCase.CreateBook(book)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed create data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success create data", books))
	}
}

func (bh *BookHandler) UpdatedBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user _entities.Book
		idToken, tokenerr := _middlewares.ReadTokenId(c)
		if tokenerr != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		}
		idn := c.Param("id")
		id, _ := strconv.Atoi(idn)
		books, rows, err := bh.bookUseCase.GetBook(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not exist"))
		}
		if idToken != int(user.UserID) {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("cannot edit this book id"))
		}
		c.Bind(&books)
		books, err = bh.bookUseCase.UpdatedBook(books, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed edit data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success edit data", books))
	}
}
