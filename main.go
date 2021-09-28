package main

import (
	"main/src/cmd/server"
	"os"
	"strconv"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	server.Run(port)
}
