package routers

import (
	"github/go-gin-example/pkg/setting"
	v1 "github/go-gin-example/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/api/v1")
	{
		// 获取标签
		apiv1.GET("/tags", v1.GetTags)
		// 新建标签
		apiv1.POST("/tag", v1.AddTag)
		// 修改标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r
}