package workflow

import (
	"context"
	"log"

	"github.com/Julien4218/workflow-er-poc/models/signals"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"go.temporal.io/sdk/client"
)

var (
	WorkflowID            string
	RunID                 string
	SlackIsIncidentSignal bool
)

var SignalCommand = &cobra.Command{
	Use:   "signal-workflow",
	Short: "Signal workflow command",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.Dial(client.Options{})
		if err != nil {
			log.Fatalf("client error: %v\n", err)
		}
		defer c.Close()

		signal := signals.SlackIsIncidentSignal{
			IsIncident: SlackIsIncidentSignal,
		}
		err = c.SignalWorkflow(context.Background(), WorkflowID, RunID, signals.SlackIsIncidentSignalName, signal)
		if err != nil {
			logrus.Fatal("Error sending the Signal")
		}

	},
}

func init() {
	SignalCommand.Flags().StringVar(&WorkflowID, "workflowID", "", "WorkflowID")
	_ = SignalCommand.MarkFlagRequired("workflowID")
	SignalCommand.Flags().StringVar(&RunID, "runID", "", "RunID")
	_ = SignalCommand.MarkFlagRequired("runID")
	SignalCommand.Flags().BoolVar(&SlackIsIncidentSignal, "isIncident", true, "IsIncident")
	_ = SignalCommand.MarkFlagRequired("isIncident")
}
