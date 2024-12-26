package cmd

import (
	"errors"
	"log"

	"github.com/fanialfi/gotodo/internal/task"
	"github.com/spf13/cobra"
)

func NewAddCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add task to the task list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAddTaskCMD(cmd)
		},
	}

	cmd.Flags().String("description", "", "name task")
	err := cmd.MarkFlagRequired("description")
	if err != nil {
		log.Printf("Error marking flags as required : %s\n", err.Error())
		return nil
	}
	return cmd
}

func runAddTaskCMD(cmd *cobra.Command) error {
	desctiption, err := cmd.Flags().GetString("description")
	if err != nil {
		return err
	}

	if len(desctiption) == 0 {
		return errors.New("description for task is required")
	}

	return task.AddTask(desctiption)
}
