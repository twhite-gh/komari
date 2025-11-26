package main

import (
	"log"

	"github.com/komari-monitor/komari/cmd"
	"github.com/komari-monitor/komari/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("Komari Monitor %s (hash: %s)", utils.CurrentVersion, utils.VersionHash)
	cmd.Execute()
}
