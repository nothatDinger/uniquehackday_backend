package handler

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"unique_hackday/db"
	"unique_hackday/errorcode"
)


// ?tag_name="美食"
func GetSceneListByTagNameHandler(c *gin.Context){
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	valid := validation.Validation{}
	//if arg := c.Query("tag_id"); arg != "" {
	//	tagID := com.StrTo(arg).MustInt()
	//	valid.Min(tagID, 1, "ID").Message("ID必须大于0")
	//	data["id"] = tagID
	//}

	if arg := c.Query("tag_name"); arg != "" {
		tagName := com.StrTo(arg).String()
		valid.Required(tagName, "tag_name").Message("必须要有tag标签用于索引")
		maps["name"] = tagName
	}

	code := errorcode.SUCCESS

	data["list"] = db.GetTagList(maps)//db.GetTagByName(maps["tagName"].(string))
	data["total"] = len(data["list"].([]db.Tag))

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : errorcode.GetErrMsg(code),
		"data" : data,
	})
}


// data = {"name":"光谷","address":"珞喻路呐呐呐","longitude":123.12,"latitude":23.21,tags":["美食","购物"]}
func AddSceneHandler(c *gin.Context) {
	var data db.Scene
	c.BindJSON(&data)

	//var tags []db.Tag
	//for _, v := range data["tags"].([]string){
	//	tags  = append(tags, db.Tag{Name: v})
	//}
	//data["tags"] = tags

	//log.Printf("data %v", data.Name)
	db.AddScenes(data)

	c.JSON(http.StatusOK,gin.H{
		"code": errorcode.SUCCESS,
		"msg": errorcode.GetErrMsg(errorcode.SUCCESS),
		"data": make(map[string]interface{}),
	})
}