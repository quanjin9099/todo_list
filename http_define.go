package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RspResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"response"`
}

func GinRsp(c *gin.Context, statusCode int, obj interface{}) {
	c.JSON(statusCode, obj)
}

func GinResultSimple(c *gin.Context, code int, msg string) {
	rsp := RspResult{}
	rsp.Code = code
	rsp.Msg = msg
	GinRsp(c, http.StatusOK, rsp)
}

func GinResultObj(c *gin.Context, code int, msg string, obj interface{}) {
	rsp := RspResult{}
	rsp.Code = code
	rsp.Msg = msg
	rsp.Data = obj
	GinRsp(c, http.StatusOK, rsp)
}
