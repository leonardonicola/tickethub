package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardonicola/tickethub/internal/router"
)

func main() {
	r := gin.Default()
	router.InitRoutes(r)
	if err := r.Run(":3000"); err != nil {
		panic(err.Error())
	}
}
