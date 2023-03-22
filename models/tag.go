package models

type Tag struct {
	TagId      int    `gorm:"primaryKey" json:"tag_id"`
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int64) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("tag_id").Where("name = ?", name).Find(&tag)
	return tag.TagId > 0
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

// 标签是否存在数据库

func ExistTagById(tag_id int) bool {
	var tag Tag
	db.Select("tag_id").Where("id = ?", tag_id).Find(&tag)
	return tag.TagId > 0
}

// 修改标签
func EditTag(tag_id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", tag_id).Updates(data)
	return true
}

// 删除标签

func DeleteTag(tag_id int) bool {
	db.Where("id = ?", tag_id).Delete(&Tag{})
	return true
}
