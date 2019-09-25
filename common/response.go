package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//相应数据
type ResponsePageData struct {
	Total int         `json:"total"`
	Items interface{} `json:"items"`
}

// 相应页面
type ResponsePage struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    ResponsePageData `json:"data"`
}

// 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	c.JSON(status, v)
	c.Abort()
}
func ResSuccessPage(c *gin.Context, total int, list interface{}) {
	ret := ResponsePage{Code: 0, Message: "ok", Data: ResponsePageData{Total: total, Items: list}}
	ResJSON(c, http.StatusOK, &ret)
}

type ResponseModelBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ResErrSrv(c *gin.Context, err error) {
	ret := ResponseModelBase{Code: 400, Message: "服务端故障， " + err.Error()}
	ResJSON(c, http.StatusBadRequest, ret)
}
