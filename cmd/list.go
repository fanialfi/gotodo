package cmd

import (
	"github.com/fanialfi/gotodo/internal/task"
	"github.com/fanialfi/gotodo/lib"
	"github.com/spf13/cobra"
)

func NewListCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runListCMD(cmd)
		},
	}

	cmd.Flags().Bool("done", false, "list task status done")
	cmd.Flags().Bool("in-progress", false, "list task status in progress")
	cmd.Flags().Bool("todo", false, "list task status todo")

	return cmd
}

func runListCMD(cmd *cobra.Command) error {
	done, err := cmd.Flags().GetBool("done")
	if err != nil {
		return err
	}

	in_progress, err := cmd.Flags().GetBool("in-progress")
	if err != nil {
		return err
	}

	todo, err := cmd.Flags().GetBool("todo")
	if err != nil {
		return err
	}

	if done {
		tasks, err := task.ListTask(task.TASK_STATUS_DONE)
		if err != nil {
			return err
		}

		lib.PrintingOutput(tasks, "Tasks (Done)")
	} else if in_progress {
		tasks, err := task.ListTask(task.TASK_STATUS_IN_PROGRESS)
		if err != nil {
			return err
		}

		lib.PrintingOutput(tasks, "Tasks (in-progress)")
	} else if todo {
		tasks, err := task.ListTask(task.TASK_STATUS_TODO)
		if err != nil {
			return err
		}

		lib.PrintingOutput(tasks, "Tasks (todo)")
	} else {
		tasks, err := task.ListTask("all")
		if err != nil {
			return err
		}

		lib.PrintingOutput(tasks, "Tasks (all)")
	}

	return nil
}
