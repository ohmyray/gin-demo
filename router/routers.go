package router

import (
"github.com/gin-gonic/gin"
"github.com/ohmyray/gin-demo01/controller"
)

// 全局前缀
const apiPreFix = "/api/v1"

func CollectRoute(r *gin.Engine) *gin.Engine {

	r.POST(apiPreFix+"/register",controller.Register)

	return r
}