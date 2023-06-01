package main

import (
	"github.com/alobe/seawill/controller"
	"github.com/alobe/seawill/lib"
)

func main() {
	lib.InitDB()
	controller.InitRouter()
	defer lib.CloseDB()
}
