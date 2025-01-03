package task

import (
	"fmt"
	"slices"
	"time"
)

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "todo"
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
		Status:      TASK_STATUS_TODO,
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

func DeleteTask(taskID int64) error {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return err
	}

	indexTask := slices.IndexFunc(tasks, func(task Task) bool {
		return task.ID == taskID
	})

	if indexTask >= 0 {
		tasks = append(tasks[:indexTask], tasks[indexTask+1:]...)
	}

	tasks = slices.Clip(tasks)

	return WriteTaskToFile(tasks)
}

func MarkInProgressTask(taskID int64) error {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return err
	}

	// create map for fast lookup based on ID
	tasksMap := make(map[int64]*Task)
	for i := range tasks {
		tasksMap[tasks[i].ID] = &tasks[i]
	}

	if task, exisi := tasksMap[taskID]; exisi {
		task.Status = TASK_STATUS_IN_PROGRESS
		task.UpdateAt = time.Now().UnixMilli()
	} else {
		return fmt.Errorf("task id : %d not found", taskID)
	}

	return WriteTaskToFile(tasks)
}

func MarkDoneTask(taskID int64) error {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return err
	}

	taskMap := make(map[int64]*Task)
	for i := range tasks {
		taskMap[tasks[i].ID] = &tasks[i]
	}

	if task, exist := taskMap[taskID]; exist {
		task.Status = TASK_STATUS_DONE
		task.UpdateAt = time.Now().UnixMilli()
	} else {
		return fmt.Errorf("task id : %d not found", taskID)
	}

	return WriteTaskToFile(tasks)
}

func ListTask(status TaskStatus) (*[]Task, error) {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return nil, err
	}

	taskMap := make(map[TaskStatus]*Task)

	switch status {
	case TASK_STATUS_DONE:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_DONE {
				taskMap[task.Status] = &task
			}
		}
	case TASK_STATUS_TODO:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_TODO {
				taskMap[task.Status] = &task
			}
		}
	case TASK_STATUS_IN_PROGRESS:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_IN_PROGRESS {
				taskMap[task.Status] = &task
			}
		}
	default:
		return &tasks, nil
	}

	tasksResult := make([]Task, 0, len(taskMap))
	for _, task := range taskMap {
		tasksResult = append(tasksResult, *task)
	}

	return &tasksResult, nil
}
