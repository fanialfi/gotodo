package cmd

import (
	"errors"

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
			return runUpdateTaskCMD(cmd)
		},
	}

	cmd.Flags().Int64("id", 0, "id for the task")
	cmd.Flags().String("description", "", "new description for the task")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("description")
	return cmd
}

func runUpdateTaskCMD(cmd *cobra.Command) error {
	id, err := cmd.Flags().GetInt64("id")
	if err != nil {
		return err
	}

	if id == 0 {
		return errors.New("id for the task is required")
	}

	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return err
	}

	if len(description) == 0 {
		return errors.New("new description for the task is required")
	}

	return task.UpdateTask(id, description)
}
