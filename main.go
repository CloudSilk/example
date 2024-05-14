package main

import (
	"flag"
	"fmt"

	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/CloudSilk/curd/gen"
	curdhttp "github.com/CloudSilk/curd/http"
	curdmodel "github.com/CloudSilk/curd/model"
	curdservice "github.com/CloudSilk/curd/service"
	"github.com/CloudSilk/pkg/constants"
	"github.com/CloudSilk/pkg/db"
	"github.com/CloudSilk/pkg/db/mysql"
	"github.com/CloudSilk/pkg/db/sqlite"
	"github.com/CloudSilk/pkg/utils"
	ucconfig "github.com/CloudSilk/usercenter/config"
	uchttp "github.com/CloudSilk/usercenter/http"
	ucmodel "github.com/CloudSilk/usercenter/model"
	"github.com/CloudSilk/usercenter/model/token"
	ucmiddleware "github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	webPath := flag.String("ui", "./web", "web路径")
	port := flag.Int("port", 80, "端口")
	flag.Parse()
	StartAll(*webPath, *port)

}

// E:\work\go\src\codeup.aliyun.com\atali\nooocode\SwiftEase\examples\complex
func StartAll(webPath string, port int) {
	err := ucconfig.InitFromFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	var dbClient db.DBClientInterface
	if ucconfig.DefaultConfig.DBType == "sqlite" {
		dbClient = sqlite.NewSqlite2("", "", ucconfig.DefaultConfig.Sqlite, "CompaAI", ucconfig.DefaultConfig.Debug)
	} else {
		dbClient = mysql.NewMysql(ucconfig.DefaultConfig.Mysql, ucconfig.DefaultConfig.Debug)
	}
	ucmodel.InitDB(dbClient, ucconfig.DefaultConfig.Debug)
	curdmodel.InitDB(dbClient, ucconfig.DefaultConfig.Debug)
	curdservice.Init()

	token.InitTokenCache(ucconfig.DefaultConfig.Token.Key, ucconfig.DefaultConfig.Token.RedisAddr, ucconfig.DefaultConfig.Token.RedisName, ucconfig.DefaultConfig.Token.RedisPwd, ucconfig.DefaultConfig.Token.Expired)
	constants.SetPlatformTenantID(ucconfig.DefaultConfig.PlatformTenantID)
	constants.SetSuperAdminRoleID(ucconfig.DefaultConfig.SuperAdminRoleID)
	constants.SetDefaultRoleID(ucconfig.DefaultConfig.DefaultRoleID)
	constants.SetEnabelTenant(ucconfig.DefaultConfig.EnableTenant)
	ucmodel.SetDefaultPwd(ucconfig.DefaultConfig.DefaultPwd)

	gen.LoadCache()

	r := gin.Default()
	r.Use(ucmiddleware.AuthRequired)
	r.Use(utils.Cors())

	uchttp.RegisterAuthRouter(r)
	curdhttp.RegisterRouter(r)
	r.Static("/web", webPath)
	r.Run(fmt.Sprintf(":%d", port))
}
