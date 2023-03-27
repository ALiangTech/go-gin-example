package models

type Articles struct {
	ArticleId int    `json:"article_id" gorm:"article_id"`
	TagID     int    `json:"tag_id" gorm:"tag_id"`
	Tag       Tag    `json:"tag"`
	Title     string `json:"title" gorm:"title"`
	Resume    string `json:"resume" gorm:"resume"`
	Content   string `json:"content" gorm:"content"`
	State     int    `json:"state" gorm:"state"`
}

func ExistArticleByID(id int) bool {
	var article Articles
	db.Select("article_id").Where("article_id = ?", id).Find(&article)
	return article.ArticleId > 0
}

func GetArticlesTotal(maps interface{}) (count int64) {
	db.Model(&Articles{}).Where(maps).Count(&count)
	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Articles) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

// 获取单个文章
func GetArticle(id int) (article Articles) {
	db.Where("article_id", id).First(&article)
	db.Model(&article)
	return
}

// 编辑文章

func EditArticle(id int, data interface{}) bool {
	db.Model(&Articles{}).Where("id = ?", id).Save(data)
	return true
}

// 添加文章

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Articles{
		TagID:   data["tag_id"].(int),
		Title:   data["title"].(string),
		Resume:  data["resume"].(string),
		Content: data["content"].(string),
		State:   data["state"].(int),
	})
	return true
}

// 删除文章

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Articles{})
	return true
}
