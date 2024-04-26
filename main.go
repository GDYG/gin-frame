package main

import (
	"gin-frame/router"
)

func main() {
	r := router.Router()

	r.Run(":9999")
}
