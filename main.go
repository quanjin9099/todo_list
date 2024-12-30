package main

// @title TodoList 项目API文档
// @version 1.0
// @description TodoList 项目API文档
// @termsOfService http://swagger.io/terms/

// @contact.name jin quan
// @contact.url
// @contact.email bothwin.qj@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 这里写接口服务的host
// @BasePath /
func main() {
	MainDb = InitDB("./main.db")
	SecDb = InitDB("./sec.db")
	DbList = append(DbList, MainDb)
	DbList = append(DbList, SecDb)
	InitServer(9081)
}
