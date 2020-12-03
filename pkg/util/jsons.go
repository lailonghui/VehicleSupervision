package util

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

type respObjList struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Total   int         `json:"total"`
	Data    interface{} `json:"data"`
}

type respObj struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type respObjError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// type GinContext struct {
// 	GinC *gin.Context
// }

// 压缩JSON
func CompressJson(data string) string {
	data = strings.Replace(data, " ", "", -1)
	data = strings.Replace(data, "\n", "", -1)
	return data
}

// 格式化JSON
func FormatJson(data string) (string, error) {
	var str bytes.Buffer
	err := json.Indent(&str, []byte(data), "", "  ")
	if err != nil {
		return "", err
	}
	return str.String(), nil
}

// 获取响应的JSON
func GetRespJson(code int, msg string, data interface{}) []byte {
	var resp respObj

	resp.Code = code
	if msg != "" {
		resp.Message = msg
	} else {
		resp.Message = "服务器错误！"
	}
	if data != nil {
		resp.Data = data
	} else {
		resp.Data = nil
	}
	resp_json, _ := json.Marshal(&resp)
	return resp_json
}

// 素材列表响应
func GetRespListJson(status int, message string, total int, data []map[string]interface{}) string {
	resp_map := make(map[string]interface{})
	resp_map["code"] = status
	resp_map["message"] = message
	resp_map["total"] = total
	resp_map["data"] = data

	b, _ := json.Marshal(resp_map)
	return string(b)
}

// 获取响应的JSON
func GetRespJsonForUser(code int, msg string, data map[string]interface{}) string {
	resp_map := make(map[string]interface{})

	resp_map["code"] = code
	resp_map["msg"] = msg
	resp_map["data"] = data

	b, _ := json.Marshal(resp_map)
	return string(b)
}

// func SetGinContext(gc *gin.Context){
// 	c = gc
// }

func GinJsonWithoutData(c *gin.Context, code int, msg string) {
	var resp respObjError

	resp.Code = code
	if msg != "" {
		resp.Message = msg
	} else {
		resp.Message = "服务器错误！"
	}

	c.JSON(code, resp)
}

func GinJson(c *gin.Context, code int, msg string, data interface{}) {
	var resp respObj

	resp.Code = code
	if msg != "" {
		resp.Message = msg
	} else {
		resp.Message = "服务器错误！"
	}
	if data != nil {
		resp.Data = data
	} else {
		resp.Data = nil
	}
	c.JSON(code, resp)
}

func GinJsonList(c *gin.Context, code int, message string, total int, data []map[string]interface{}) {
	var respList respObjList

	respList.Code = code
	respList.Message = message
	respList.Total = total
	respList.Data = data

	c.JSON(code, respList)
}
