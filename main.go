package main

import (
	"github.com/captainkovalsky/stream/cmd"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}
func main() {
	cmd.Execute()
}
