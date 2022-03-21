package model

import (
	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

// Task型はuuid.UUID型のID、文字列のNameとbool値のFinishedをパラメーターとして持つ
type Task struct {
	ID       uuid.UUID
	Name     string
	Finished bool
}

// 関数GetTasksは、引数はなく、戻り値は[]Task型（Task型のスライス）とerror型である
func GetTasks() ([]Task, error) {

	// 空のタスクのスライスである、tasksを定義する
	var tasks []Task

	// tasksにDBのタスク全てを代入する。その操作の成否をerrと定義する(*5)
	err := db.Find(&tasks).Error

	// tasksとerrを返す
	return tasks, err
}
