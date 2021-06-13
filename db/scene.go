package db

type Scene struct {
	Model
	Name string `json:"name"`
	Address string `json:"address"`
	Longitude float64 `json:"longitude"`//经度
	Latitude float64 `json:"latitude"`//纬度
	Tags []Tag	`json:"tags" gorm:"many2many:scene_tags"`
}

func GetSceneList(maps interface{}) (scenes []Scene) {
	db.Preload("Tags").Where(maps).Find(&scenes)
	return
}

func GetSceneByID(id int)(scene Scene) {
	db.Preload("Tags").Where("id = ?", id).First(&scene)
	return
}

func GetSceneByName(name string) (scene Scene) {
	db.Preload("Tags").Where("name = ?", name).First(&scene)
	return
}

func GetSceneTotal(maps interface{}) (count int64) {
	db.Preload("Tags").Model(&Scene{}).Where(maps).Count(&count)
	return
}

func AddScenes(data Scene) bool {
	db.Create(&Scene{
		Name: data.Name,
		Tags: data.Tags,
	})
	return true
}