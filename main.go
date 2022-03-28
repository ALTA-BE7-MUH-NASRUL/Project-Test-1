package main

import (
	"fmt"
	"latihan/coba-project/config"
	_authHandler "latihan/coba-project/delivery/handler/auth"
	_bookHandler "latihan/coba-project/delivery/handler/book"
	_loanHandler "latihan/coba-project/delivery/handler/loan"
	_returHandler "latihan/coba-project/delivery/handler/retur"
	_userHandler "latihan/coba-project/delivery/handler/user"
	_middlewares "latihan/coba-project/delivery/middleware"
	_routes "latihan/coba-project/delivery/route"
	_authRepository "latihan/coba-project/repo/auth"
	_bookRepository "latihan/coba-project/repo/book"
	_loanRepository "latihan/coba-project/repo/loan"
	_returRepository "latihan/coba-project/repo/retur"
	_userRepository "latihan/coba-project/repo/user"
	_authUseCase "latihan/coba-project/usecase/auth"
	_bookUseCase "latihan/coba-project/usecase/book"
	_loanUseCase "latihan/coba-project/usecase/loan"
	_returUseCase "latihan/coba-project/usecase/retur"
	_userUseCase "latihan/coba-project/usecase/user"
	_utils "latihan/coba-project/utils"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.GetConfig()
	db := _utils.InitDB(config)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)
	bookRepo := _bookRepository.NewBookRepository(db)
	bookUseCase := _bookUseCase.NewBookUseCase(bookRepo)
	bookHandler := _bookHandler.NewBookHandler(bookUseCase)
	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)
	loanRepo := _loanRepository.NewLoanRepository(db)
	loanUseCase := _loanUseCase.NewLoanUseCase(loanRepo)
	loanHandler := _loanHandler.NewLoanHandler(loanUseCase)
	returRepo := _returRepository.NewReturRepository(db)
	returUseCase := _returUseCase.NewReturUseCase(returRepo)
	returHandler := _returHandler.NewReturHandler(returUseCase)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middlewares.CustomLogger())
	_routes.RegisterPath(e, userHandler, bookHandler, loanHandler, returHandler)
	_routes.RegisterAuthPath(e, authHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
