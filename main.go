package main

import (
	"SWTAC/config"
	"SWTAC/datasource"
	"SWTAC/datasource/ent"
	"SWTAC/log"
	"context"
	"fmt"
	"github.com/bykof/gostradamus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
)

func main() {

	fmt.Println("Hello, world!!!")

	zapLogger, _ := zap.NewProduction()
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())
	e.Use(log.ZapLogger(zapLogger))
	e.IPExtractor = echo.ExtractIPDirect()
	e.IPExtractor = echo.ExtractIPFromXFFHeader()
	e.IPExtractor = echo.ExtractIPFromRealIPHeader()
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(30)))

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// 连接数据库
	client, err := datasource.NewClient()
	if err != nil {
		zapLogger.Fatal("opening ent client", zap.Error(err))
		return
	}
	// 关闭数据库
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)
	//fmt.Println(client)
	dateTime := gostradamus.Now()
	fmt.Println(dateTime)

	//defer client.Close()
	ctx := context.Background()

	//autoMigration := database.AutoMigration
	autoMigration := datasource.AutoMigration
	autoMigration(client, ctx)

	//debugMode := database.DebugMode
	debugMode := datasource.DebugMode
	debugMode(err, client, ctx)

	e.Group("user")

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world!!!")
	})

	// user
	//e.GET("/users", handler.GetAllUsers(client))
	//e.GET("/user/id", handler.GetUserById(client))
	//e.GET("/user/name", handler.GetUserByUsername(client))
	//e.GET("/user/email", handler.GetUserByEmail(client))
	//e.GET("/user/phone", handler.GetUserByPhone(client))
	//e.POST("/user", handler.CreateUser(client))
	//e.PUT("/user/id", handler.UpdateUserById(client))
	//e.PUT("/user", handler.UpdateUser(client))
	//e.DELETE("/user", handler.DeleteUser(client))
	//e.DELETE("/user/id", handler.DeleteUserById(client))
	//// usergroup
	//e.GET("/usergroups", handler.GetAllUserGroups(client))
	//e.GET("/usergroup/id", handler.GetUserGroupById(client))
	//e.GET("/usergroup/name", handler.GetUserGroupByName(client))
	//e.POST("/usergroup", handler.CreateUserGroup(client))
	//e.PUT("/usergroup/id", handler.UpdateUserGroupById(client))
	//e.PUT("/usergroup", handler.UpdateUserGroup(client))
	//e.DELETE("/usergroup", handler.DeleteUserGroup(client))
	//e.DELETE("/usergroup/id", handler.DeleteUserGroupById(client))
	//
	//// role
	//e.GET("/roles", handler.GetAllRoles(client))
	//e.GET("/role/id", handler.GetRoleById(client))
	//e.GET("/role/name", handler.GetRoleByRoleName(client))
	//e.POST("/role", handler.CreateRole(client))
	//e.PUT("/role/id", handler.UpdateRoleById(client))
	//e.PUT("/role", handler.UpdateRole(client))
	//e.DELETE("/role", handler.DeleteRole(client))
	//e.DELETE("/role/id", handler.DeleteRoleById(client))

	e.Logger.Fatal(e.Start(":2023"))

}

func init() {
	config.InitConfig()
}
