package main

import (
	"os"

	"github.com/pipeproxy/pipe-xds/internal/cmd/proxy"
	"github.com/pipeproxy/pipe-xds/internal/cmd/request"
	"github.com/pipeproxy/pipe-xds/internal/cmd/wait"
	"github.com/spf13/cobra"
	"istio.io/istio/pkg/cmd"
	cleaniptables "istio.io/istio/tools/istio-clean-iptables/pkg/cmd"
	iptables "istio.io/istio/tools/istio-iptables/pkg/cmd"
	"istio.io/pkg/log"
)

var (
	loggingOptions = log.DefaultOptions()

	rootCmd = &cobra.Command{
		Use:          "pipe-xds",
		Short:        "Istiod Agent.",
		Long:         "Istiod Agent runs in the sidecar or gateway container and bootstraps Pipe.",
		SilenceUsage: true,
		FParseErrWhitelist: cobra.FParseErrWhitelist{
			// Allow unknown flags for backward-compatibility.
			UnknownFlags: true,
		},
	}
)

func init() {
	// Attach the Istio logging options to the command.
	loggingOptions.AttachCobraFlags(rootCmd)

	cmd.AddFlags(rootCmd)

	rootCmd.AddCommand(proxy.GetCommand())
	rootCmd.AddCommand(iptables.GetCommand())
	rootCmd.AddCommand(cleaniptables.GetCommand())
	rootCmd.AddCommand(wait.GetCommand())
	rootCmd.AddCommand(request.GetCommand())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Errora(err)
		os.Exit(-1)
	}
}
