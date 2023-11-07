package main

import (
	"flag"
	"log"

	"github.com/spf13/cobra"

	"github.com/Julien4218/temporal-signal-workflow/activity"
	"github.com/Julien4218/temporal-signal-workflow/workflow"
)

// Command represents the base command when called without any subcommands
var Command = &cobra.Command{
	PersistentPreRun:  globalInit,
	Use:               "temporal-send",
	Short:             "The CLI to trigger and signal workflow",
	Long:              `The CLI allows to trigger new workflow instance or signal existing workflow activities`,
	DisableAutoGenTag: true, // Do not print generation date on documentation
}

func init() {
	// Bind imported sub-commands
	Command.AddCommand(workflow.Command)
	Command.AddCommand(activity.Command)
	Command.AddCommand(workflow.SignalCommand)
}

func globalInit(cmd *cobra.Command, args []string) {
	// Initialize logger
	// logLevel := configAPI.GetLogLevel()
	// config.InitLogger(log.StandardLogger(), logLevel)

	// Initialize client
	// if client.NRClient == nil {
	// 	client.NRClient = createClient()
	// }
}

func init() {
	// Command.PersistentFlags().BoolVar(&config.FlagDebug, "debug", false, "debug level logging")
}

func main() {
	if err := Command.Execute(); err != nil {
		if err != flag.ErrHelp {
			log.Fatal(err)
		}
	}
}
