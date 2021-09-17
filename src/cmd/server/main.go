package server

import (
	"fmt"
	"main/src/app"
)

func Run(port int) {

	if err := app.Engine.Run(fmt.Sprintf(":%d",port)); err != nil {
		panic(err)
	}
}