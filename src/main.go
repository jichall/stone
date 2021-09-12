package main

import (
	"flag"

	"go.uber.org/zap"
)

var (
	host = flag.String("host", "localhost", "host of the server")
	port = flag.String("port", "8080", "port of the server")

	logger *zap.Logger
)

func main() {
	flag.Parse()

	// The error variable declaration is not necessary, but it is used to
	// prevent variable shadowing.
	var err error
	logger, err = zap.NewProduction()

	if err != nil {
		panic(err)
	}

	serve(*host, *port)
}
