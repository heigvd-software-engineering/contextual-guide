package main

import (
	"flag"
	"main/src/cmd/server"
)

func main() {
	port := flag.Int("port",3000, "-port=3000")
	flag.Parse()

	server.Run(*port)
}
