package commands

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "url-monitor",
	Short: "url-monitor is used monitor api urls",
	Long:  `url-monitor is used monitor api urls and some more detail stuff along with it`,
}

/*
The init function is responsible to run things
which we will require before anything else
say
  - Fetch API Keys
  - Set Logging level
  - Setup any environment variable required for the app
*/
func init() {
	rootCmd.AddCommand(checkUrlCmd, checkStatusCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
