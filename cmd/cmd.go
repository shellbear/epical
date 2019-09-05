package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	EpitechToken string
	Credentials  string
)

var cli = &cobra.Command{
	Use:   "epical",
	Short: "Synchronize your Epitech's events with Google calendar",
	Long: "A fast and simple way to sync your Epitech calendar with Google.\n" +
		"Complete documentation is available at https://github.com/shellbear/epical",
}

func Execute() {
	cli.PersistentFlags().StringVarP(&Credentials, "credentials", "c", "credentials.json", "Google API credentials")

	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	listCmd.PersistentFlags().StringVarP(&EpitechToken, "token", "t", "", "Epitech API Token")
	listCmd.MarkPersistentFlagRequired("token")
	syncCmd.PersistentFlags().StringVarP(&EpitechToken, "token", "t", "", "Epitech API Token")
	syncCmd.MarkPersistentFlagRequired("token")

	cli.AddCommand(versionCmd)
	cli.AddCommand(clearCmd)
	cli.AddCommand(listCmd)
	cli.AddCommand(syncCmd)
}
