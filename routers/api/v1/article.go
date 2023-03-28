package v1

import (
	"github/go-gin-example/models"
	"github/go-gin-example/pkg/e"
	"github/go-gin-example/pkg/setting"
	"github/go-gin-example/pkg/util"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 获取单个文章
// 肯定是根据文章id 来获取
// 第一性原则
func GetArticle(ctx *gin.Context) {
	article_id := com.StrTo(ctx.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(article_id, 1, "id").Message("ID不符合规范")
	code := e.INVALID_PARAMS

	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(article_id) {
			data = models.GetArticle(article_id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 获取多个文章
func GetArticles(ctx *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1

	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(rune(state)).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只允许0或1 ")
	}
	var tagId int = -1
	if arg := ctx.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		valid.Min(tagId, 0, "tag_id").Message("标签ID必须大于0")
	}
	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		data["list"] = models.GetArticles(util.GetPage(ctx), setting.PageSize, maps)
		data["total"] = models.GetArticlesTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.Key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 新增文章

func AddArticle(ctx *gin.Context) {
	tagId := com.StrTo(ctx.Query("tag_id")).MustInt()
	title := ctx.Query("title")
	resume := ctx.Query("resume")
	content := ctx.Query("content")
	state := com.StrTo(ctx.DefaultQuery("state", "0")).MustInt()
	valid := validation.Validation{}
	valid.Min(tagId, 0, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(resume, "resume").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只运行0或q")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistTagById(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["resume"] = resume
			data["content"] = content
			data["state"] = state

			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Panicf("err.Key: %s, err.Message: %s", err.Key, err.Message)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

// 修改文章

func EditArticle(ctx *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(ctx.Param("id")).MustInt()
	tagId := com.StrTo(ctx.Query("tag_id")).MustInt()
	title := ctx.Query("title")
	resume := ctx.Query("resume")
	content := ctx.Query("content")
	modifiedBy := ctx.Query("modified_by")

	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或者1")
	}
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			if models.ExistTagById(tagId) {
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if resume != "" {
					data["resume"] = resume
				}
				if content != "" {
					data["content"] = content
				}
				data["modified_by"] = modifiedBy
				models.EditArticle(id, data)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_TAG

			}
		} else {
			for _, err := range valid.Errors {
				log.Printf("err.Key: %s, err.Message: %s", err.Key, err.Message)
			}
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

// 删除文章

func DeleteArticle(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID 必须大于0")
	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.Key:%s, err.Message:%s", err.Key, err.Message)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}
