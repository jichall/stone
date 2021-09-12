package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/jichall/stone/src/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	// as I'm not using any kind of storage database, I will define variables to
	// represent every entity in my application
	accounts map[int]*models.Account = make(map[int]*models.Account)
	transactions map[int]*models.Transaction = make(map[int]*models.Transaction)
)


func serve(host, port string) {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/accounts", fetchAccounts)
	e.GET("/accounts/:id/balance", fetchBalance)
	e.POST("/accounts", createAccount)

	e.POST("/login", login)

	e.GET("/transfers", fetchTransfers)
	e.POST("/transfers", createTransfer)

	// define a custom server using the host and port given
	e.Server = &http.Server{
		Addr:         strings.Join([]string{host, port}, ":"),
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	e.Logger.Fatal(e.Start(e.Server.Addr))
}
