package main

import (
	. "gin-frame/dao"
	"gin-frame/router"
)

func main() {
	r := router.Router()

	InitializeDB()

	r.Run(":9999")
}
