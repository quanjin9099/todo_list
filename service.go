package main

import (
	"fmt"
	"github.com/TelenLiu/knife4j_go"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
	"sort"
	"todo_list/docs"
)

func HealthTest(c *gin.Context) {
	c.String(200, `ok`)
}

type TodoListRequest struct {
	UserId int `json:"user_id" form:"user_id" `
	Limit  int `json:"limit" form:"limit"`
}

type TodoInfoResponse struct {
	Id          int    `json:"id"`          //ID
	Name        string `json:"name"`        //名称
	Description string `json:"description"` //描述
	Score       int    `json:"score"`       //分数 1-100
}

type ToDoListResponse struct {
	UserName string             `json:"user_name"`
	Data     []TodoInfoResponse `json:"data"`
}

type TodoResponse struct {
	Code     int              `json:"code"`
	Message  string           `json:"message"`
	Response ToDoListResponse `json:"response"`
}

func TodoInfo2TodoInfoResponse(dataList []TodoInfo) []TodoInfoResponse {
	rList := make([]TodoInfoResponse, 0)
	for _, v := range dataList {
		rList = append(rList, TodoInfoResponse{v.Id, v.Name, v.Description, v.Score})
	}
	return rList
}

// GetUserToDoList
// @Summary 获取用户待办列表接口
// @Description
// @Tags TODOList
// @Accept application/json
// @Produce application/json
// @Param object query TodoListRequest false "列表参数"
// @Success 200 {object} TodoResponse
// @Router /user/todo_list [get]
func GetUserToDoList(c *gin.Context) {
	var req TodoListRequest
	if err := c.ShouldBind(&req); err != nil {
		GinResultSimple(c, 0x0001, "请求参数错误:"+err.Error())
		return
	}

	DataList := make([]TodoInfo, 0)
	for _, v := range DbList {
		tmpList := make([]TodoInfo, 0)
		if err := v.Model(&TodoInfo{}).Where("user_id=?", req.UserId).Limit(req.Limit).Scan(&tmpList).Error; err != nil {
			fmt.Printf("GetUserToDoList error! %+v\n", err)
			continue
		} else {
			if len(tmpList) == 0 { // 没查到数据，继续下一个数据库连接
				continue
			}
			// 查到数据了
			DataList = append(DataList, tmpList...) // 合并数据
			sort.Sort(TodoInfoSlice(DataList))      // 排序
			DataList = DataList[:req.Limit]         // 截取长度
		}
	}
	var userInfo User
	if err := DbList[0].Where("id=?", req.UserId).First(&userInfo).Error; err != nil {
		fmt.Printf("GetUserToDoList error! %+v\n", err)
		GinResultSimple(c, 0x0002, "用户不存在")
		return
	}
	result := TodoResponse{
		Code:     0,
		Message:  "查询成功",
		Response: ToDoListResponse{UserName: userInfo.Name, Data: TodoInfo2TodoInfoResponse(DataList)},
	}
	GinRsp(c, http.StatusOK, result)
}

func InitServer(port int) {
	r := gin.Default()
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	//把swag 生成的 docs.go 或 swagger.json 写入
	knife4j_go.SetDiyContent("doc.json", []byte(docs.SwaggerInfo.ReadDoc()))
	//挂载knife4j主题ui到 你要目录，如 /doc
	r.StaticFS("/doc", http.FS(knife4j_go.GetKnife4jVueDistRoot()))
	r.GET("/services.json", func(c *gin.Context) {
		c.String(200, `[
		    {
				"name": "ToDoList",
				"url": "/doc.json",
				"swaggerVersion": "2.0",
				"location": "/doc.json",
			}
		]`)
	})
	r.Use(CORSMiddleware())
	r.GET("/health_check", HealthTest)
	r.GET("/user/todo_list", GetUserToDoList)
	runAddr := fmt.Sprintf("0.0.0.0:%v", port)
	fmt.Printf("InitServer>%s", runAddr)
	if err := r.Run(runAddr); err != nil {
		fmt.Printf("InitServer error! %+v", err.Error())
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE,PATCH")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Content-Type, XFILENAME,XFILECATEGORY,XFILESIZE,x-requested-with, Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == http.MethodOptions {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
