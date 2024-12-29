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
			return runListCMD()
		},
	}

	return cmd
}

func runListCMD() error {
	tasks, err := task.ListTask()
	if err != nil {
		return err
	}

	fmt.Println("Tasks (All)")
	for _, task := range *tasks {
		creationTime := time.UnixMilli(task.CreatedAt).Format("Monday, 02-Jan-2006 15:00:4")
		fmt.Printf("ID:%d\t%s\t\t%s (%v)\n", task.ID, task.Status, task.Description, creationTime)
	}

	return nil
}
