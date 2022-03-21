package router

// 使用するライブラリをimport
import (
	"net/http"

	"github.com/ToDoList-Server/model"

	"github.com/labstack/echo/v4"
)

// 関数 GetTasksHandlerは引数がecho.Context型のc で、戻り値はerror型である
func GetTasksHandler(c echo.Context) error {

	// model(package)の関数GetTasksを実行し、戻り値をtasks,errと定義する。
	tasks, err := model.GetTasks()

	// errが空でない時は StatusBadRequest(*5) を返す
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	// StasusOK と tasksを返す
	return c.JSON(http.StatusOK, tasks)
}
