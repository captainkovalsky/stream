package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the server",
	Long:  `Run the http server which emits cyclic json response`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Server streamer")
	},
}
