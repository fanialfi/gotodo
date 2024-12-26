package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gotodo",
		Short: "gotodo is a CLI tool for managing tasks",
		Long: `gotodo is a CLI tool for managing tasks. It allow you to create, list, edit, and delete task
		
you can also mark task as completed and update their status
complete code available at https://github.com/fanialfi/gotodo`,
	}

	cmd.AddCommand(NewAddCMD())
	cmd.AddCommand(NewUpdateCMD())
	cmd.AddCommand(NewDeleteCMD())
	cmd.AddCommand(NewMarkInProgressCMD())
	return cmd
}
