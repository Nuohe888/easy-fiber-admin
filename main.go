// @title EasyFiberAdmin API
// @version 1.0
// @description This is the API documentation for EasyFiberAdmin.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:18888
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description "Type 'Bearer YOUR_JWT_TOKEN' to authenticate"
package main

import "easy-fiber-admin/boot"

func main() {
	// 启动
	boot.Boot()
}
