/*
@Time : 2020/12/3 16:32
@Author : lai
@Description : 返回函数，可自定义
@File : response
*/
package util

import (
	"VehicleSupervision/pkg/status"
	"github.com/gin-gonic/gin"
	"net/http"
)

//response函数
func commonResponse(c *gin.Context, err error, data ...interface{}) {
	if err != nil {
		GinJsonWithoutData(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(data) == 0 {
		GinJsonWithoutData(c, status.StatusOK, status.StatusText(status.StatusOK))
	} else if len(data) == 1 {
		GinJson(c, status.StatusOK, status.StatusText(status.StatusOK), nil)
	}
}
