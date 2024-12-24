package task

import (
	"time"
)

type TaskStatus string

const (
	TASK_STATUS_DONE        TaskStatus = "done"
	TASK_STATUS_IN_PROGRESS TaskStatus = "in-progress"
)

type Task struct {
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	ID          int64      `json:"id"`
	CreatedAt   int64      `json:"created-at"`
	UpdateAt    int64      `json:"updated-at"`
}

func newTask(taskId int64, desctiption string) *Task {
	return &Task{
		Description: desctiption,
		Status:      TASK_STATUS_IN_PROGRESS,
		ID:          taskId,
		CreatedAt:   time.Now().UnixMilli(),
		UpdateAt:    time.Now().UnixMilli(),
	}
}

func AddTask(descrition string) error {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return err
	}

	var newTaskId int64
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		newTaskId = lastTask.ID + 1
	} else {
		newTaskId = 1
	}

	task := newTask(newTaskId, descrition)
	tasks = append(tasks, *task)

	return WriteTaskToFile(tasks)
}
