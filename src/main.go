package main

import (
	"flag"
)

var (
	host = flag.String("host", "localhost", "host of the server")
	port = flag.String("port", "8080", "port of the server")
)

func main() {
	flag.Parse()
	serve(*host, *port)
}
