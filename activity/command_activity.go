package activity

import (
	"context"
	"encoding/base64"
	"log"
	"path"

	"github.com/spf13/cobra"
	"go.temporal.io/sdk/client"
)

var (
	Token  string
	Signal string
	Body   string
)

// signalCmd represents the signal command
var Command = &cobra.Command{
	Use:   "signal",
	Short: "signal an activity",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.Dial(client.Options{})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}
		defer c.Close()

		// Get required workflow inputs
		wfid, runid, err := WorkflowFromToken(Token)
		if err != nil {
			log.Fatalf("%s", err.Error())
			return
		}

		log.Printf("Got workflow ID:%s, runid:%s", wfid, runid)

		// Move signalName to param
		err = c.SignalWorkflow(
			context.Background(),
			wfid,
			runid,
			Signal,
			Body,
		)
		if err != nil {
			log.Fatalf("%s", err.Error())
			return
		}

		log.Println("done")
	},
}

func init() {
	Command.Flags().StringVar(&Token, "token", "", "Token")
	_ = Command.MarkFlagRequired("token")
	Command.Flags().StringVar(&Body, "body", "", "Body")
	_ = Command.MarkFlagRequired("body")
	Command.Flags().StringVar(&Signal, "signal", "", "Signal")
	_ = Command.MarkFlagRequired("signal")
}

func WorkflowFromToken(token string) (string, string, error) {
	var rawToken []byte

	rawToken, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return "", "", err
	}

	wfid := path.Dir(string(rawToken))
	runid := path.Base(string(rawToken))

	return wfid, runid, nil
}
