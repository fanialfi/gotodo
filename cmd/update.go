package cmd

import (
	"errors"
	"strconv"

	"github.com/fanialfi/gotodo/internal/task"
	"github.com/spf13/cobra"
)

func NewUpdateCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update sebuah task",
		Long: `update sebuah task dengan memberikan task ID dan new status
		Example
		gotodo update 1 'new description'`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUpdateTaskCMD(args)
		},
	}
	return cmd
}

func runUpdateTaskCMD(args []string) error {
	if len(args) == 0 {
		return errors.New("task id & task description is required")
	}

	taskID := args[0]
	taskIDInt, err := strconv.ParseInt(taskID, 10, 0)
	if err != nil {
		return err
	}
	description := args[1]

	return task.UpdateTask(taskIDInt, description)
}
