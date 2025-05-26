package main

import (
	"github.com/gin-gonic/gin"

	"backend/middleware"
	"backend/route"

	_ "backend/init"
)

func main() {
	r := gin.Default()

	middleware.RegisterMiddleware(r)
	route.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
