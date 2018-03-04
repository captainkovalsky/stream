package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clientCmd)
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run the client",
	Long:  `Run the http stream reader client`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Client stream reader")
	},
}
