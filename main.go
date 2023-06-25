package main

import (
	"SWTAC/config"
	"SWTAC/datasource"
	"SWTAC/datasource/ent"
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
	log, _ := zap.NewDevelopment()
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.IPExtractor = echo.ExtractIPDirect()
	e.IPExtractor = echo.ExtractIPFromXFFHeader()
	e.IPExtractor = echo.ExtractIPFromRealIPHeader()
	e.Use(middleware.Logger())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// 连接数据库

	client, err := datasource.NewClient()
	if err != nil {
		log.Fatal("opening ent client", zap.Error(err))
		return
	}
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

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world!!!")
	})

	e.Logger.Fatal(e.Start(":2023"))

}

func init() {
	config.InitConfig()
}
