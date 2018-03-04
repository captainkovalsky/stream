package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "App",
	Long:  ``,
}

func init() {
	viper.SetDefault("author", "Viktor Dzundza <victor.dzundza.dev@gmail.com>")
}

//Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.WithField("err", err).Panic("Can not execute root command")
		os.Exit(1)
	}
}
