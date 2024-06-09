package main

import (
	"flag"
	"fmt"

	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/CloudSilk/curd/gen"
	curdhttp "github.com/CloudSilk/curd/http"
	curdmodel "github.com/CloudSilk/curd/model"
	curdservice "github.com/CloudSilk/curd/service"
	"github.com/CloudSilk/pkg/config"
	"github.com/CloudSilk/pkg/constants"
	"github.com/CloudSilk/pkg/utils"
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

func StartAll(webPath string, port int) {
	err := config.InitFromFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	ok, dbClient := config.NewDB("sqlite")
	if !ok {
		panic("未配置数据库")
	}
	ucmodel.InitDB(dbClient, config.DefaultConfig.Debug)
	curdmodel.InitDB(dbClient, config.DefaultConfig.Debug)
	curdservice.Init()

	token.InitTokenCache(config.DefaultConfig.Token.Key, config.DefaultConfig.Token.RedisAddr, config.DefaultConfig.Token.RedisName, config.DefaultConfig.Token.RedisPwd, config.DefaultConfig.Token.Expired)
	constants.SetPlatformTenantID(config.DefaultConfig.PlatformTenantID)
	constants.SetSuperAdminRoleID(config.DefaultConfig.SuperAdminRoleID)
	constants.SetDefaultRoleID(config.DefaultConfig.DefaultRoleID)
	constants.SetEnabelTenant(config.DefaultConfig.EnableTenant)
	ucmodel.SetDefaultPwd(config.DefaultConfig.DefaultPwd)

	gen.LoadCache()

	r := gin.Default()
	r.Use(ucmiddleware.AuthRequired)
	r.Use(utils.Cors())

	uchttp.RegisterAuthRouter(r)
	curdhttp.RegisterRouter(r)
	r.Static("/web", webPath)
	r.Run(fmt.Sprintf(":%d", port))
}
