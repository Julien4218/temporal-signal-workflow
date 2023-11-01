package workflow

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/xtgo/uuid"
	"go.temporal.io/sdk/client"
)

var (
	Name  string
	Input string
)

var Command = &cobra.Command{
	Use:   "workflow",
	Short: "Workflow command",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.Dial(client.Options{})
		if err != nil {
			log.Fatalf("client error: %v\n", err)
		}
		defer c.Close()

		log.Printf("Getting input")
		rawInput := "{}"
		if len(Input) > 0 {
			rawInput, err := base64.URLEncoding.DecodeString(Input)
			if err != nil {
				log.Fatal(err)
				return
			}
			log.Printf("Got rawInput:[%s]\n", rawInput)
		}
		var jsonInput interface{}
		err = json.Unmarshal([]byte(rawInput), &jsonInput)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Got jsonInput:%s\n", jsonInput)

		workflowName := Name
		workflowID := fmt.Sprintf("%s-%s", workflowName, uuid.NewRandom().String())
		queueName := fmt.Sprintf("%s-Queue", Name)
		options := client.StartWorkflowOptions{
			ID:        workflowID,
			TaskQueue: queueName,
		}
		log.Printf("Starting workflow ID:%s queue:%s\n", workflowID, queueName)
		we, err := c.ExecuteWorkflow(context.Background(), options, Name, jsonInput)
		if err != nil {
			log.Fatalf("Unable to start the workflow ID:%s queue:%s error:%s\n", workflowID, queueName, err.Error())
		}

		log.Printf("workflow ID:%s runID:%s\n", we.GetID(), we.GetRunID())
		log.Println("done")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {
	Command.Flags().StringVar(&Name, "name", "", "Name")
	_ = Command.MarkFlagRequired("name")
	Command.Flags().StringVar(&Input, "input", "", "input")
}
