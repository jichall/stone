package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func serve(host, port string) {

	e := echo.New()

	e.Use(middleware.Logger())

	// define a custom server using the host and port given
	e.Server = &http.Server{
		Addr:         strings.Join([]string{host, port}, ":"),
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	e.Logger.Fatal(e.Start(e.Server.Addr))
}
