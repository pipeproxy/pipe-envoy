package request

import (
	"net/http"
	"time"

	"github.com/spf13/cobra"

	"istio.io/istio/pilot/pkg/request"
)

// NB: extra standard output in addition to what's returned from envoy
// must not be added in this command. Otherwise, it'd break istioctl proxy-config,
// which interprets the output literally as json document.
var (
	requestCmd = &cobra.Command{
		Use:   "request <method> <path> [<body>]",
		Short: "Makes an HTTP request to the Envoy admin API",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(c *cobra.Command, args []string) error {
			command := &request.Command{
				Address: "localhost:15000",
				Client: &http.Client{
					Timeout: 60 * time.Second,
				},
			}
			body := ""
			if len(args) >= 3 {
				body = args[2]
			}
			return command.Do(args[0], args[1], body)
		},
	}
)

func GetCommand() *cobra.Command {
	return requestCmd
}
