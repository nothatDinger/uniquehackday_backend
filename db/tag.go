package db

type Tag struct {
	Model
	Name string `json:"name"`

	Scenes []Scene `json:"scenes" gorm:"many2many:scene_tags"`
}

func GetTagList(maps interface{}) (tags []Tag) {
	db.Preload("Scenes").Where(maps).Find(&tags)
	return
}

func GetTagByID(id int) (tag []Tag) {
	db.Preload("Scenes").Where("id = ?", id).Find(&tag)
	return
}

func GetTagByName(name string) (tag []Tag) {
	db.Preload("Scenes").Where("name = ?", name).Find(&tag)
	return
}

func GetTagTotal(maps interface{}) (count int64) {
	db.Preload("Scenes").Where(maps).Count(&count)
	return
}

func AddTag(data []Tag) bool {
	db.Create(&data)
	return true
}
