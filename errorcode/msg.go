package errorcode

import "log"

var MsgFlags = map[int]string{
	SUCCESS:                        "OK",
	ERROR:                          "FAIL",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "标签名已存在",
	ERROR_NOT_EXIST_TAG:            "标签名不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章已存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已过期",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "无效的Token",
}

func GetErrMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok{
		return msg
	}
	log.Print("There are something wrong when getting error code message")
	return MsgFlags[ERROR]
}