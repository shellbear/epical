package cmd

import (
	"fmt"
	"github.com/ShellBear/epical/pkg/epical"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync all upcoming Epitech events to Google calendar",
	Run: func(cmd *cobra.Command, args []string) {
		epical.SyncCalendar(cmd.Flag("credentials").Value.String(), cmd.Flag("token").Value.String())
		fmt.Println("Successfully synchronized Epitech events to Google calendar.")
	},
}
