package cmd

import (
	"github.com/ShellBear/epical/pkg/epical"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all upcoming Epitech events",
	Run: func(cmd *cobra.Command, args []string) {
		epical.ListEvents(cmd.Flag("token").Value.String())
	},
}
