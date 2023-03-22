package v1

import (
	"fmt"
	"github/go-gin-example/models"
	"github/go-gin-example/pkg/e"
	"github/go-gin-example/pkg/setting"
	"github/go-gin-example/pkg/util"
	"net/http"

	"github.com/astaxie/beego/validation"
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
	name := ctx.Query("name")
	state := com.StrTo(ctx.DefaultQuery("state", "0")).MustInt()
	createdBy := ctx.Query("created_by")
	fmt.Print(createdBy)
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或者1")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		// 校验通过
		if !models.ExistTagByName(name) {
			// 不存在相同标签名称
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 修改标签

func EditTag(ctx *gin.Context) {
	tag_id := com.StrTo(ctx.Param("tag_id")).MustInt()
	name := ctx.Query("name")
	state := com.StrTo(ctx.DefaultQuery("state", "-1")).MustInt()
	valid := validation.Validation{}

	valid.Range(state, 0, 1, "state").Message("状态只允许0")
	valid.Required(tag_id, "tag_id").Message("tag_id不能为空")
	valid.Required(name, "name").Message("标签名称不能为空")
	valid.MaxSize(name, 100, "name").Message("标签名称最长为100字符")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		// 没有错误
		code = e.SUCCESS
		if models.ExistTagById(tag_id) {
			// 存在才可以修改
			data := make(map[string]interface{})
			data["name"] = name
			data["state"] = state
			models.EditTag(tag_id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 删除标签
// 标签是否存在
// 存在就删除
// 否则告知不存在

func DeleteTag(ctx *gin.Context) {
	tag_id := com.StrTo(ctx.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(tag_id, 1, "tag_id").Message("ID必须大于0")
	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagById(tag_id) {
			models.DeleteTag(tag_id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
