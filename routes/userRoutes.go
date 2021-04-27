package routes

import (
    controller "go_user_login/controllers"
    "github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
    incomingRoutes.POST("/user/signup", controller.SignUp())
    incomingRoutes.POST("/user/login", controller.Login())
}
