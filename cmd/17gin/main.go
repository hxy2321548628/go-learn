package main

import (
	"fmt"
	"go-learn/internal/middleware/auth"
	"go-learn/internal/middleware/cors"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// const MaxSize = 5<<10 // 1kb 1024 bit
const MaxSize = 10 << 20 // 5mb

type User struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age"`
}

func main() {
	router := gin.Default()
	router.Use(cors.CorsMdwe)

	userGroup := router.Group("/user")
	userGroup.Use(auth.AuthMdwe)

	router.GET("/", func(ctx *gin.Context) {
		fmt.Println("你好")
		ctx.String(http.StatusOK, "你好")
	})

	// 路径参数：/user/12
	userGroup.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.String(http.StatusOK, "id=%s", id)
	})

	// Query：/user?id=1&name=小满
	userGroup.GET("", func(ctx *gin.Context) {
		id := ctx.DefaultQuery("id", "none")
		name := ctx.DefaultQuery("name", "none")
		ctx.String(http.StatusOK, "id=%s name=%s", id, name)
	})

	// 表单正文：id=1&name=小满
	userGroup.POST("/formdata", func(ctx *gin.Context) {
		id := ctx.DefaultPostForm("id", "none")
		name := ctx.DefaultPostForm("name", "none")
		ctx.String(http.StatusOK, "id=%s name=%s", id, name)
	})

	// JSON 正文：{"name":"小满","age":18}
	userGroup.POST("/create", func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"code":    http.StatusCreated,
			"data":    user,
			"success": true,
		})
	})

	// 单文件上传
	router.POST("/single-upload", func(ctx *gin.Context) {
		// 获取文件
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		// 可能的校验
		if file.Size > MaxSize {
			ctx.String(http.StatusBadRequest, "文件过大")
			return
		}

		// 处理文件名
		dst := filepath.Join("./files/", filepath.Base(file.Filename))

		// 保存文件
		if err := ctx.SaveUploadedFile(file, dst); err != nil {
			ctx.String(http.StatusBadRequest, "上传失败，请重试")
			return
		}
		ctx.String(http.StatusOK, "保存成功")

	})

	// 多文件上传
	router.POST("/multi-upload", func(ctx *gin.Context) {
		// 多文件上传， 文件会放在 MultipartForm 中
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		files := form.File["files"]
		if len(files) == 0 {
			ctx.String(http.StatusBadRequest, "没有文件")
			return
		}

		for _, file := range files {

			// 可能的校验
			if file.Size > MaxSize {
				ctx.String(http.StatusBadRequest, "文件过大")
				return
			}

			// 处理文件名
			dst := filepath.Join("./files/", filepath.Base(file.Filename))

			// 保存文件
			if err := ctx.SaveUploadedFile(file, dst); err != nil {
				ctx.String(http.StatusBadRequest, "上传失败，请重试")
				return
			}
		}
		ctx.String(http.StatusOK, "保存成功")

	})

	if err := router.Run(); err != nil {
		panic(err)
	}
}
