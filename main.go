package main

import (
	"github.com/ToDoList-Server/model"
	"github.com/ToDoList-Server/router"
	"github.com/labstack/echo/v4"
)

func main() {
	sqlDB := model.DBConnection()
	defer sqlDB.Close()
	e := echo.New()
	router.SetRouter(e)
}
