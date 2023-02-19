package commands

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/url"
)

var checkStatusCmd = &cobra.Command{
	Use:   "check-status",
	Short: "check-status",
	Long:  `check-status`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			log.Fatal("subcommand check-url only take one argument as url")
		}
		if _, err := url.ParseRequestURI(args[0]); err != nil {
			log.Fatalf("wrong url type [%s]", err)
		}
		err := checkStatus(args)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func checkStatus(args []string) error {
	fmt.Printf("HI!! From check-status sub-command with %s as argument", args[0])
	return nil
}
