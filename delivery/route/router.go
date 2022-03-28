package route

import (
	_authHandler "latihan/coba-project/delivery/handler/auth"
	_bookHandler "latihan/coba-project/delivery/handler/book"
	_loanHandler "latihan/coba-project/delivery/handler/loan"
	_returHandler "latihan/coba-project/delivery/handler/retur"
	_userHandler "latihan/coba-project/delivery/handler/user"
	_middlewares "latihan/coba-project/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uh *_userHandler.UserHandler, bh *_bookHandler.BookHandler, lh *_loanHandler.LoanHandler, rh *_returHandler.ReturHandler) {
	// user
	e.GET("/users", uh.GetAllHandler(), _middlewares.JWTMiddleware())
	e.POST("/users", uh.CreateUserHandler())
	e.GET("/users/:id", uh.GetUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())
	e.PUT("/users/:id", uh.UpdatedUserHandler(), _middlewares.JWTMiddleware())

	// book
	e.GET("/books", bh.GetAllHandler())
	e.POST("/books", bh.CreateBookHandler(), _middlewares.JWTMiddleware())
	e.GET("/books/:id", bh.GetBookHandler())
	e.DELETE("/books/:id", bh.DeleteBookHandler(), _middlewares.JWTMiddleware())
	e.PUT("/books/:id", bh.UpdatedBookHandler(), _middlewares.JWTMiddleware())

	e.POST("/loan", lh.LoaningHandler())
	e.POST("/retur", rh.ReturingHandler())
}

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}
