package router

// 使用するライブラリをimport
import (
	"net/http"

	"github.com/ToDoList-Server/model"
	"github.com/google/uuid"

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

// ReqTask型は文字列のNameをパラメーターとして持つ
type ReqTask struct {
	Name string `json:"name"`
}

// 関数 AddTaskHandler は引数がecho.Context型で、戻り値はerror型である
func AddTaskHandler(c echo.Context) error {

	// 空のReqTaskである、reqを定義
	var req ReqTask

	// bodyのjsonファイルをbind(*1)
	err := c.Bind(&req)
	// エラーハンドリング(day-3のGetTaskHandlerを参照)
	// StatusBadRequestを返す
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	// 空のmodel(package)のTaskである、taskを定義
	var task *model.Task

	// model(package)のAddTask関数を実行し、戻り値をtask,errと定義
	task, err = model.AddTask(req.Name)
	// エラーハンドリング
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	// StastsOK と 追加されたtaskを返す
	return c.JSON(http.StatusOK, task)
}

func ChangeFinishedTaskHandler(c echo.Context) error {

	// taskIDのパスパラメータ(string型)を取得し、uuid型に変換。その値をtaskID、成否をerrとする
	taskID, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	// 関数 ChangeFinishedTaskを実行、戻り値をerrに代入する(errを更新した)
	err = model.ChangeFinishedTask(taskID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.NoContent(http.StatusOK)
}

func DeleteTaskHandler(c echo.Context) error {
	taskID, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	err = model.DeleteTask(taskID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.NoContent(http.StatusOK)
}
