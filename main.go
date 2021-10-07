package main

import (
	"github.com/gopla/maro/src/model"
	"github.com/gopla/maro/src/router"
)

func main() {
	model.ConnectDataBase()

	r:= router.SetupRouter()

	r.Run()
}