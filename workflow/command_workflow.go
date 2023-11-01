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
		rawInput, err := getBase64Decode(Input)
		if err != nil {
			// input is not base64 encoded
			rawInput = Input
		}
		param, err := getJsonDecode(rawInput)
		if err != nil {
			log.Fatalf("Invalid json input receieved:%s detail:%s\n", rawInput, err.Error())
		}
		log.Printf("Got jsonInput:%s\n", param)

		workflowName := Name
		workflowID := fmt.Sprintf("%s-%s", workflowName, uuid.NewRandom().String())
		queueName := fmt.Sprintf("%s-Queue", Name)
		options := client.StartWorkflowOptions{
			ID:        workflowID,
			TaskQueue: queueName,
		}
		log.Printf("Starting workflow ID:%s queue:%s\n", workflowID, queueName)
		we, err := c.ExecuteWorkflow(context.Background(), options, Name, param)
		if err != nil {
			log.Fatalf("Unable to start the workflow ID:%s queue:%s error:%s\n", workflowID, queueName, err.Error())
		}

		log.Printf("workflow ID:%s runID:%s\n", we.GetID(), we.GetRunID())
		log.Println("done")
	},
}

func getBase64Decode(input string) (string, error) {
	if len(input) > 0 {
		rawInput, err := base64.URLEncoding.DecodeString(input)
		if err != nil {
			return "", err
		}
		return string(rawInput), nil
	}
	return "", nil
}

func getJsonDecode(input string) (interface{}, error) {
	var result interface{}
	if len(input) > 0 {
		err := json.Unmarshal([]byte(input), &result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {
	Command.Flags().StringVar(&Name, "name", "", "Name")
	_ = Command.MarkFlagRequired("name")
	Command.Flags().StringVar(&Input, "input", "", "input")
}
