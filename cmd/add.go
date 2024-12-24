package cmd

import (
	"errors"

	"github.com/fanialfi/gotodo/internal/task"
	"github.com/spf13/cobra"
)

func NewAddCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add task to the task list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAddTaskCMD(args)
		},
	}
	return cmd
}

func runAddTaskCMD(args []string) error {
	if len(args) == 0 {
		return errors.New("task description is required")
	}
	desctiption := args[0]
	return task.AddTask(desctiption)
}
