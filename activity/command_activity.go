package activity

import (
	"context"
	"log"
	"path"

	"github.com/spf13/cobra"

	"github.com/Julien4218/temporal-signal-workflow/temporal"
	"github.com/Julien4218/temporal-signal-workflow/util"
)

var (
	Token  string
	Signal string
	Input  string
)

func init() {
	Command.Flags().StringVar(&Token, "token", "", "Token")
	_ = Command.MarkFlagRequired("token")

	Command.Flags().StringVar(&Signal, "signal", "", "Signal")
	_ = Command.MarkFlagRequired("signal")

	Command.Flags().StringVar(&Input, "input", "", "Input")
}

// signalCmd represents the signal command
var Command = &cobra.Command{
	Use:   "signal",
	Short: "signal an activity",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := temporal.GetTemporalClient()
		if err != nil {
			log.Fatalf("client error: %v\n", err)
		}
		defer c.Close()

		// Get required workflow inputs
		wfid, runid, err := WorkflowFromToken(Token)
		if err != nil {
			log.Fatalf("%s", err.Error())
			return
		}

		log.Printf("Got workflow ID:%s, runid:%s", wfid, runid)

		param := util.GetInputParam(Input)

		// Move signalName to param
		err = c.SignalWorkflow(
			context.Background(),
			wfid,
			runid,
			Signal,
			param,
		)
		if err != nil {
			log.Fatalf("%s", err.Error())
			return
		}

		log.Println("done")
	},
}

func WorkflowFromToken(token string) (string, string, error) {
	rawToken, err := util.GetBase64Decode(token)
	if err != nil {
		return "", "", err
	}

	wfid := path.Dir(string(rawToken))
	runid := path.Base(string(rawToken))

	return wfid, runid, nil
}
