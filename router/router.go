package router

import (
	"os"

	"github.com/labstack/echo/v4/middleware"

	_ "net/http"

	"github.com/labstack/echo/v4"
)

//Routingを設定する関数　引数はecho.echo型であり、戻り値はerror型
func SetRouter(e *echo.Echo) error {

	// 諸々の設定(*1)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header}\n",
		Output: os.Stdout,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// APIを書く場所
	e.GET("/api/tasks", GetTasksHandler)

	// 8000番のポートを開く(*2)
	err := e.Start(":8000")
	return err
}
