package cmd

import (
	"log"

	"github.com/spf13/cobra"
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
	cli.PersistentFlags().StringVarP(&Credentials, "credentials", "c", "./", "Google API credentials folder")

	if err := cli.Execute(); err != nil {
		log.Fatalln("Failed to execute CLI", err)
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
