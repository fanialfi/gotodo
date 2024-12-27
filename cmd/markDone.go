package cmd

import (
	"errors"

	"github.com/fanialfi/gotodo/internal/task"
	"github.com/spf13/cobra"
)

func NewMarkDoneCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-done",
		Short: "mark done a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runMarkDoneCMD(cmd)
		},
	}

	cmd.Flags().Int64("id", 0, "id for the task")
	cmd.MarkFlagRequired("id")

	return cmd
}

func runMarkDoneCMD(cmd *cobra.Command) error {
	id, err := cmd.Flags().GetInt64("id")
	if err != nil {
		return err
	}

	if id == 0 {
		return errors.New("id for the task is required")
	}

	return task.MarkDoneTask(id)
}
