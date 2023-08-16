package response

import (
	"github.com/gin-gonic/gin"
	"Leilei-platform/public/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

const (
	ERROR   = 1
	SUCCESS = 0
)

func (g *Gin) Result(httpCode, errCode int, data interface{}) {
	// 开始时间
	g.C.JSON(httpCode, Response{
		Status: errCode,
		Msg:    e.GetMsg(errCode),
		Data:   data,
	})
}
