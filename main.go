package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ohmyray/gin-demo01/common"
	"github.com/ohmyray/gin-demo01/router"
)


func main() {
	// 获取数据库连接
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = router.CollectRoute(r)

	panic(r.Run())
}
