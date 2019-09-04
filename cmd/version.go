package cmd

import (
	"fmt"
	"github.com/ShellBear/epical/pkg/epical"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Epical",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Epical v%s\n", epical.VERSION)
	},
}
