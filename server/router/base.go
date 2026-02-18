package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("base")
	baseRouter.POST("Captcha", api.ApiGroupApp.Captcha)
	baseRouter.POST("SendEmailVerificationCode", api.ApiGroupApp.SendEmailVerificationCode)
	baseRouter.GET("QQLoginURL", api.ApiGroupApp.QQLoginURL)
}
