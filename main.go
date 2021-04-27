package main

import (
    "os"
    middleware "go_user_login/middleware"
    routes "go_user_login/routes"
    "github.com/gin-gonic/gin"
    _ "github.com/heroku/x/hmetrics/onload"
)

func main() {
    port := os.Getenv("PORT")

    if port == "" {
        port = "8000"
    }

    router := gin.New()
    router.Use(gin.Logger())
    routes.UserRoutes(router)

    router.Use(middleware.Authentication())

    router.GET("/dummy-api", func(c *gin.Context) {

        c.JSON(200, gin.H{"success": "Access granted for dummy API."})

    })

    router.Run(":" + port)
}