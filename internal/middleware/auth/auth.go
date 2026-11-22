package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
鉴权中间件
*/
func AuthMdwe(ctx *gin.Context) {
	fmt.Println("=====> auth before")
	ctx.Next()
	fmt.Println("=====> auth end")
}
