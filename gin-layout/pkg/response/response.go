package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Page struct {
	List      interface{} `json:"list"`
	Count     int         `json:"count"`
	PageIndex int         `json:"page_index"`
	PageSize  int         `json:"page_size"`
}

type PageResponse struct {
	Code int    `json:"code"`
	Data Page   `json:"data"`
	Msg  string `json:"msg"`
}

func (res *Response) ReturnError(code int, msg string) {
	res.Code = http.StatusInternalServerError
	res.Message = "系统异常"

	if code != 0 {
		res.Code = code
	}
	if msg != "" {
		res.Message = msg
	}
}

func (res *Response) ReturnOK(msg string, data interface{}) {
	res.Code = 0
	res.Message = "请求成功"
	res.Data = data

	if msg != "" {
		res.Message = msg
	}
}

func (res *PageResponse) ReturnOK(list interface{}, count, pageIndex, pageSize int, msg string) {
	res.Code = 0
	res.Data.List = list
	res.Data.Count = count
	res.Data.PageIndex = pageIndex
	res.Data.PageSize = pageSize
	res.Msg = "请求成功"

	if msg != "" {
		res.Msg = msg
	}
}

func Error(c *gin.Context, code int, msg string) {
	var res Response
	res.ReturnError(code, msg)
	c.JSON(http.StatusOK, res)
}

func OK(c *gin.Context, data interface{}, msg string) {
	var res Response
	res.ReturnOK(msg, data)

	c.JSON(http.StatusOK, res)
}

func PageOK(c *gin.Context, list interface{}, count, pageIndex, pageSize int, msg string) {
	var res PageResponse
	res.ReturnOK(list, count, pageIndex, pageSize, msg)

	c.JSON(http.StatusOK, res)
}
