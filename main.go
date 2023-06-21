package main

import (
	"sfcup/db"
	"sfcup/router"
)

func main() {
	db.InitGorm()
	//db.InitGen()
	router.EngineStart()
}
