package workflow

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/xtgo/uuid"
	"go.temporal.io/sdk/client"

	"github.com/Julien4218/temporal-signal-workflow/temporal"
	"github.com/Julien4218/temporal-signal-workflow/util"
)

var (
	WorkflowID   string
	WorkflowType string
	QueueName    string
	Input        string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {
	Command.Flags().StringVar(&WorkflowID, "workflowID", "", "WorkflowID")

	Command.Flags().StringVar(&WorkflowType, "workflowType", "", "WorkflowType")
	_ = Command.MarkFlagRequired("workflowType")

	Command.Flags().StringVar(&QueueName, "queue", "", "Queue")
	_ = Command.MarkFlagRequired("queue")

	Command.Flags().StringVar(&Input, "input", "", "Input")
}

var Command = &cobra.Command{
	Use:   "workflow",
	Short: "Workflow command",
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			fmt.Printf("Error loading .env file")
			os.Exit(1)
		}

		c, err := temporal.GetTemporalClient()
		if err != nil {
			log.Fatalf("client error: %v\n", err)
		}
		defer c.Close()

		param := util.GetInputParam(Input)

		if WorkflowID == "" {
			WorkflowID = uuid.NewRandom().String()
		}

		options := client.StartWorkflowOptions{
			ID:        WorkflowID,
			TaskQueue: QueueName,
		}
		log.Printf("Starting workflow ID:%s type:%s queue:%s\n", WorkflowID, WorkflowType, QueueName)
		we, err := c.ExecuteWorkflow(context.Background(), options, WorkflowType, param)
		if err != nil {
			log.Fatalf("Unable to start the workflow ID:%s queue:%s error:%s\n", WorkflowID, QueueName, err.Error())
		}

		log.Printf("workflow ID:%s runID:%s\n", we.GetID(), we.GetRunID())
		log.Println("done")
	},
}
