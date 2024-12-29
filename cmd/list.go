package cmd

import (
	"fmt"
	"time"

	"github.com/fanialfi/gotodo/internal/task"
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

		fmt.Println("Tasks (Done)")
		for _, task := range *tasks {
			creationTime := time.UnixMilli(task.CreatedAt).Format("Monday, 02-Jan-2006 15:00:4")
			fmt.Printf("ID:%d\t%s\t\t%s (%v)\n", task.ID, task.Status, task.Description, creationTime)
		}
	} else if in_progress {
		tasks, err := task.ListTask(task.TASK_STATUS_IN_PROGRESS)
		if err != nil {
			return err
		}

		fmt.Println("Tasks (in-progress)")
		for _, task := range *tasks {
			creationTime := time.UnixMilli(task.CreatedAt).Format("Monday, 02-Jan-2006 15:00:4")
			fmt.Printf("ID:%d\t%s\t\t%s (%v)\n", task.ID, task.Status, task.Description, creationTime)
		}
	} else if todo {
		tasks, err := task.ListTask(task.TASK_STATUS_TODO)
		if err != nil {
			return err
		}

		fmt.Println("Tasks (todo)")
		for _, task := range *tasks {
			creationTime := time.UnixMilli(task.CreatedAt).Format("Monday, 02-Jan-2006 15:00:4")
			fmt.Printf("ID:%d\t%s\t\t%s (%v)\n", task.ID, task.Status, task.Description, creationTime)
		}
	} else {
		tasks, err := task.ListTask("all")
		if err != nil {
			return err
		}

		fmt.Println("Tasks (all)")
		for _, task := range *tasks {
			creationTime := time.UnixMilli(task.CreatedAt).Format("Monday, 02-Jan-2006 15:00:4")
			fmt.Printf("ID:%d\t%s\t\t%s (%v)\n", task.ID, task.Status, task.Description, creationTime)
		}
	}

	return nil
}
