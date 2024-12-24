package task

import (
	"fmt"
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

	fmt.Printf("Task added successfully (ID:%d)\n\n", newTaskId)
	return WriteTaskToFile(tasks)
}

func UpdateTask(taskId int64, description string) error {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return err
	}

	// create map for fast lookup based on ID
	tasksMap := make(map[int64]*Task)
	for i := range tasks {
		tasksMap[tasks[i].ID] = &tasks[i]
	}

	if task, exisi := tasksMap[taskId]; exisi {
		task.Description = description
		task.UpdateAt = time.Now().UnixMilli()
	} else {
		return fmt.Errorf("task id : %d not found", taskId)
	}

	return WriteTaskToFile(tasks)
}
