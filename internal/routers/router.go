package routers

import (
	v1 "github.com/eighteen20/blog-service/internal/routers/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	article := v1.NewArticle()
	tag := v1.NewTag()

	apiV1 := r.Group("/api/v1")

	{
		apiV1.POST("tags", tag.Create)        // 新增标签
		apiV1.DELETE("/tags/:id", tag.Delete) // 删除指定标签
		apiV1.PUT("/tags/:id", tag.Update)    // 修改某标签
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags", tag.List) // 获取标签

		apiV1.POST("/articles", article.Create)       // 新增文章
		apiV1.DELETE("/articles/:id", article.Delete) // 删除指定文章
		apiV1.PUT("/articles/:id", article.Update)    // 更新指定文章
		apiV1.PATCH("/articles/:id/state", article.Update)
		apiV1.GET("/articles/:id", article.Get) // 获取指定文章
		apiV1.GET("/articles", article.List)    // 获取文章列表
	}

	return r
}
