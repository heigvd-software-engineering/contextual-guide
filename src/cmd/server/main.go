package main

import (
	"flag"
	"fmt"
	"main/src/internal/app"
)

func main() {
	port := flag.Int("port",3000, "-port=3000")
	flag.Parse()

	if err := app.Engine.Run(fmt.Sprintf(":%d",*port)); err != nil {
		panic(err)
	}
}