package main

import (
	"github.com/ToDoList-Server/model"
	"github.com/ToDoList-Server/router"
	"github.com/labstack/echo/v4"
)

func main() {
	db := model.CreateDB()
	defer db.Close()
	e := echo.New()
	router.SetRouter(e)
}
