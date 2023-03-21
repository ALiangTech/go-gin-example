package v1

import (
	"github/go-gin-example/models"
	"github/go-gin-example/pkg/e"
	"github/go-gin-example/pkg/setting"
	"github/go-gin-example/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 获取全部标签
func GetTags(ctx *gin.Context) {
	name := ctx.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS
	data["lists"] = models.GetTags(util.GetPage(ctx), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 新增文章标签
func AddTag(ctx *gin.Context) {

}

// 修改标签

func EditTag(ctx *gin.Context) {

}

// 删除标签

func DeleteTag(c *gin.Context) {

}
