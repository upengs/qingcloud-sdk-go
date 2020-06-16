package options

import (
	"github.com/spf13/cobra"
	"github.com/yunify/qingcloud-sdk-go/config"
)

var (
	cfg, err        = config.NewDefault()
	loginPasswd     = ""
	loginKeyPair    = ""
	loginMode       = "passwd"
	needNewSID      = 1
	instances       = ""
	accessKeyID     = ""
	secretAccessKey = ""
	zone            = ""
	logLevel        = ""
)

func flags(command *cobra.Command) {
	command.Flags().StringVar(&zone, "zone", "gd2", "区域 ID，注意要小写")
	command.Flags().StringVar(&accessKeyID, "access-key-id", "", "申请的 API密钥ID，例如”QYACCESSKEYIDEXAMPLE”。")
	command.Flags().StringVar(&secretAccessKey, "secret-access-key", "", "")
	command.Flags().StringVar(&logLevel, "log-lever", "debug", "")
}

func assignment(cfg *config.Config) {
	cfg.AccessKeyID = accessKeyID
	cfg.SecretAccessKey = secretAccessKey
	cfg.LogLevel = logLevel
	cfg.Zone = zone
}
