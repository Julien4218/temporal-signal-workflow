package cmd

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
	"github.com/xtgo/uuid"
	"go.temporal.io/sdk/client"
)

// startCmd represents the start workflow command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a background check",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.Dial(client.Options{
			// MetricsHandler: tallyhandler.NewMetricsHandler(newPrometheusScope(prometheus.Configuration{
			// 	ListenAddress: "0.0.0.0:8001",
			// 	TimerType:     "histogram",
			// })),
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}
		defer c.Close()

		id := uuid.NewRandom()

		// Move this to required param
		const MoneyTransferTaskQueueName = "TRANSFER_MONEY_TASK_QUEUE"
		options := client.StartWorkflowOptions{
			ID:        "pay-invoice-" + id.String(),
			TaskQueue: MoneyTransferTaskQueueName,
		}

		log.Printf("Starting transfer")

		rawInput, err := base64.URLEncoding.DecodeString(Input)
		if err != nil {
			return
		}
		log.Printf("Got rawInput:[%s]", rawInput)
		var jsonInput interface{}
		err = json.Unmarshal([]byte(rawInput), &jsonInput)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Got jsonInput:%s", jsonInput)

		we, err := c.ExecuteWorkflow(context.Background(), options, Name, jsonInput)
		if err != nil {
			log.Fatalln("Unable to start the Workflow:", err)
		}

		log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

		log.Println("done")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVar(&Name, "name", "", "Name")
	startCmd.MarkFlagRequired("name")
	startCmd.Flags().StringVar(&Input, "input", "", "input")
	startCmd.MarkFlagRequired("input")
}
