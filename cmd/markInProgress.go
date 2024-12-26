package cmd

import (
	"errors"
	"log"

	"github.com/fanialfi/gotodo/internal/task"
	"github.com/spf13/cobra"
)

func NewMarkInProgressCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-in-progress",
		Short: "mark a task as in progress",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runMarkInProgressCMD(cmd)
		},
	}

	cmd.Flags().Int("id", 0, "id for the task")
	err := cmd.MarkFlagRequired("id")
	if err != nil {
		log.Printf("Error marking flags as required : %s\n", err.Error())
		return nil
	}
	return cmd
}

func runMarkInProgressCMD(cmd *cobra.Command) error {
	id, err := cmd.Flags().GetInt("id")
	if err != nil {
		return err
	}

	if id == 0 {
		return errors.New("id for the task is required")
	}

	return task.MarkInProgressTask(int64(id))
}
