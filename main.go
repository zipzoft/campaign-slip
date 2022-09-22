package main

import (
	"campiagn-slip/config"
	"campiagn-slip/middleware"
	"campiagn-slip/pkg/database"
	"campiagn-slip/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	router := gin.Default()

	router.Use(middleware.CORS())

	routers.Routes(router)
	database.MongoConnection()
	err := router.Run(":" + strconv.Itoa(config.GetConfig().Port))

	if err != nil {
		fmt.Println(err)
		return
	}
}
