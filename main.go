package main

import (
	"main/src/cmd/server"
	"net/url"
	"os"
	"strconv"
)

func main() {
	u, err := url.Parse(os.Getenv("APP_URL"))
	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(u.Port())
	if err != nil {
		panic(err)
	}

	server.Run(port)
}
