package server

import (
	"fmt"
	"main/src/internal"
)

func Run(port int) {
	if err := internal.Engine.Run(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}
