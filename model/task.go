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

// 関数 AddTask は引数がstring型のnameで、戻り値はTaskのポインターとerror型である (*1)
func AddTask(name string) (*Task, error) {

	// 新たなuuidを生成し、これをid、成否をerrとする（*2）
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	// ID,Name,Finishedにid,name,false を代入したTask型のtaskを定義
	task := Task{
		ID:       id,
		Name:     name,
		Finished: false,
	}

	// taskをDBのTaskテーブルに追加。その成否を(ry
	err = db.Create(&task).Error

	// taskのポインタ と errを返す
	return &task, err
}

// 関数 ChangeFinishedTaskの引数はuuid.UUID型のtaskIDで、戻り値はerror型である
func ChangeFinishedTask(taskID uuid.UUID) error {

	// DBのTaskテーブルからtaskIDと一致するidを探し、そのFinishedをtureにする(*3)
	err := db.Model(&Task{}).Where("id = ?", taskID).Update("finished", true).Error
	return err
}

func DeleteTask(taskID uuid.UUID) error {
	// DBのTaskテーブルからtaskIDと一致するidを探し、そのタスクを削除する
	err := db.Where("id = ?", taskID).Delete(&Task{}).Error
	return err
}
