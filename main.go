package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ohmyray/gin-demo01/common"
	"github.com/ohmyray/gin-demo01/config"
	"github.com/ohmyray/gin-demo01/router"
	"github.com/spf13/viper"
	"log"
)

func main() {

	// 获取配置信息
	config.InitConfig()

	// 获取数据库连接
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = router.CollectRoute(r)

	serverPort := viper.GetString("server.port")
	log.Printf("Server run at http://localhost:%v", serverPort)
	if serverPort != "" {
		panic(r.Run(":" + serverPort))
	}
	panic(r.Run())
}
