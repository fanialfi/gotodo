package cmd

import (
	"errors"
	"strconv"

	"github.com/fanialfi/gotodo/internal/task"
	"github.com/spf13/cobra"
)

func NewDeleteCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "hapus task dari data",
		Long:  "hapus task dari data berdasarkan id",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDeleteTaskCMD(args)
		},
	}

	return cmd
}

func runDeleteTaskCMD(args []string) error {
	if len(args) == 0 {
		return errors.New("task id is required")
	}

	taskID := args[0]
	taskIDInt, err := strconv.ParseInt(taskID, 10, 0)
	if err != nil {
		return err
	}

	return task.DeleteTask(taskIDInt)
}
