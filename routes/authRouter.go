package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/nvlannasik/user-services-go/controllers"
)

func AuthRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("auth/login", controller.Login())
	incomingRoutes.POST("auth/register", controller.Register())
}
