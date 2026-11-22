package cors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
跨域中间件
*/
func CorsMdwe(ctx *gin.Context) {
	fmt.Println("=====> cors before")

	// 修改响应头中的参数，允许浏览器读取跨域响应
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,PATCH")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type")

	// 预检不是业务请求，回 204 后结束，不再进入 /ping
	// 预检没有对应的业务路由。中间件不接住 OPTIONS 的话，Gin 回 404，浏览器认为不允许跨域，第二次 DELETE 也不会发，终端里看不到 result。
	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	ctx.Next()
	fmt.Println("=====> cors end")
}
