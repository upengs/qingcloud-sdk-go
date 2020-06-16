package main

import (
	"flag"

	"github.com/yunify/qingcloud-sdk-go/cmd/options"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"os"
)

func main() {
	command := &cobra.Command{
		Use:   "qingyun",
		Short: "qy",
		Long:  `操作青云主机`,
		Example: `
qingyun -h
qingyun reset -h
`,
		SilenceUsage: true,
	}
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	command.AddCommand(options.MakeKingFunResetCommand())
	command.AddCommand(options.MakeAttachKeyPairs())
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}

}
