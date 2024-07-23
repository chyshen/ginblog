package main

import (
	_ "github.com/chyshen/ginblog/models"
	"github.com/chyshen/ginblog/routers"
	"github.com/chyshen/ginblog/utils"
)

//	@title			个人博客API
//	@version		1.0
//	@description	个人博客
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	scy
//	@contact.url	http://github.com/chyshen
//	@contact.email	mr.scy@outlook.com

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

//	@host		localhost:8989
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				在输入框中输入请求头中需要的Jwt授权Token，格式：Bearer Token

//	@securityDefinitions.oauth2.application	OAuth2Application
//	@tokenUrl								https://example.com/oauth/token
//	@scope.write							Grants write access
//	@scope.admin							Grants read and write access to administrative information

//	@securityDefinitions.oauth2.implicit	OAuth2Implicit
//	@authorizationUrl						https://example.com/oauth/authorize
//	@scope.write							Grants write access
//	@scope.admin							Grants read and write access to administrative information

//	@securityDefinitions.oauth2.password	OAuth2Password
//	@tokenUrl								https://example.com/oauth/token
//	@scope.read								Grants read access
//	@scope.write							Grants write access
//	@scope.admin							Grants read and write access to administrative information

// @securityDefinitions.oauth2.accessCode	OAuth2AccessCode
// @tokenUrl								https://example.com/oauth/token
// @authorizationUrl						https://example.com/oauth/authorize
// @scope.admin							Grants read and write access to administrative information

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	router := routers.NewRouter()
	err := router.Run(utils.Vcf.GetString("server.http_port"))
	if err != nil {
		panic(err)
	}
}
