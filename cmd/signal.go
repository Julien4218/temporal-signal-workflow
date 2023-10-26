package cmd

import (
	"context"
	"encoding/base64"
	"log"
	"path"

	"github.com/spf13/cobra"
	"go.temporal.io/sdk/client"
)

// signalCmd represents the signal command
var signalCmd = &cobra.Command{
	Use:   "signal",
	Short: "signal a background check",
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

		// Get required workflow inputs
		wfid, runid, err := WorkflowFromToken(Token)
		if err != nil {
			log.Fatalf("%s", err.Error())
			return
		}

		log.Printf("Got wfid:%s, runid:%s", wfid, runid)

		// Move signalName to param
		err = c.SignalWorkflow(
			context.Background(),
			wfid,
			runid,
			"signal-submission",
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
	rootCmd.AddCommand(signalCmd)
	signalCmd.Flags().StringVar(&Token, "token", "", "Token")
	signalCmd.MarkFlagRequired("token")
	signalCmd.Flags().StringVar(&Body, "body", "", "Body")
	signalCmd.MarkFlagRequired("body")

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
