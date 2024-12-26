package cmd

import (
	"errors"

	"github.com/fanialfi/gotodo/internal/task"
	"github.com/spf13/cobra"
)

func NewDeleteCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "hapus task dari data",
		Long:  "hapus task dari data berdasarkan id",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDeleteTaskCMD(cmd)
		},
	}

	cmd.Flags().Int("id", 0, "task id")
	cmd.MarkFlagRequired("id")

	return cmd
}

func runDeleteTaskCMD(cmd *cobra.Command) error {
	taskId, err := cmd.Flags().GetInt("id")
	if err != nil {
		return err
	}
	if taskId == 0 {
		return errors.New("task id is required")
	}

	return task.DeleteTask(int64(taskId))
}
