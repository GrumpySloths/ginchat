package main

import (
	"ginchat/pool"
	"ginchat/router"
	"ginchat/utility"
)

func main() {
	utility.Config_init()
	utility.Redis_init()
	utility.DataBase_init()
	pool.Pool_init()
	r := router.GetRouter()

	r.Run() // listen and serve on 0.0.0.0:8080
}
