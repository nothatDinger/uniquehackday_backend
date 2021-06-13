package handler

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"unique_hackday/db"
	"unique_hackday/errorcode"
)

// ?scene_name = "光谷"
func GetTagListBySceneNameHandler(c *gin.Context)  {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	valid := validation.Validation{}

	if arg := c.Query("scene_name") ; arg != nil{
		sceneName := com.StrTo(arg).String()
		valid.Required(sceneName,"scene_name").Message("?scene_name=what")
		valid.MaxSize(sceneName,100,"scene_name").Message("名称不超过100个字符")
		maps["scene_name"] = sceneName
	}

	data["list"] = db.GetSceneList(maps)
	data["total"] = len(data["list"].([]db.Scene))

	c.JSON(http.StatusOK,gin.H{
		"code" : errorcode.SUCCESS,
		"msg" : errorcode.GetErrMsg(errorcode.SUCCESS),
		"data" : data,
	})
}

// data = {["name":"美食","tags":["name":"光谷","name":"楚河汉街"]]}
func AddTagHandler(c *gin.Context) {
	var data []db.Tag
	c.BindJSON(&data)

	//var tags []db.Tag
	//for _, v := range data["tags"].([]string){
	//	tags  = append(tags, db.Tag{Name: v})
	//}
	//data["tags"] = tags

	//log.Printf("data %v", data.Name)
	db.AddTag(data)

	c.JSON(http.StatusOK,gin.H{
		"code": errorcode.SUCCESS,
		"msg": errorcode.GetErrMsg(errorcode.SUCCESS),
		"data": make(map[string]interface{}),
	})
}
