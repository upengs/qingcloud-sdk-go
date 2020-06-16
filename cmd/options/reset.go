package options

import (
	"errors"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/yunify/qingcloud-sdk-go/config"
	"github.com/yunify/qingcloud-sdk-go/service"
)

// MakeKingFunResetCommand ...
func MakeKingFunResetCommand() *cobra.Command {
	var (
		resetInstancesInput = &service.ResetInstancesInput{LoginMode: &loginMode, LoginPasswd: &loginPasswd, LoginKeyPair: &loginKeyPair, NeedNewSID: &needNewSID}

		err error
	)
	command := &cobra.Command{
		Use:  "reset",
		Long: `reset`,
		Example: `qingyun reset -h
                  重置qingyun node
`,
		SilenceUsage: true,
	}

	if err != nil {
		log.Error().Err(err).Send()
		return command
	}

	command.Flags().StringVar(&instances, "instances", "", "一个或多个主机ID,eg:i-ijkjfdfs,i-fdfsf")
	command.Flags().StringVar(&loginPasswd, "login-passwd", loginPasswd, "登录密码")
	command.Flags().StringVar(&loginKeyPair, "login-key-pair", "", "登录密钥ID")
	command.Flags().StringVar(&loginMode, "login-mode", loginMode, "指定登录方式。当为 linux 主机时，有效值为 keypair 和 passwd; 当为 windows 主机时，只能选用 passwd 登录方式。\n当登录方式为 keypair 时，需要指定 login_keypair 参数；\n当登录方式为 passwd 时，需要指定 login_passwd 参数。")
	command.Flags().IntVar(&needNewSID, "need-new-sid", needNewSID, "1: 生成新的SID，0: 不生成新的SID, 默认为0；只对Windows类型主机有效。")
	flags(command)
	command.Run = func(c *cobra.Command, _ []string) {
		assignment(cfg)
		if err := run(cfg, resetInstancesInput); err != nil {
			log.Error().Err(err).Str("instances", instances).Msg("reset qingyun node failed")
			os.Exit(1)
		}
	}
	return command
}

func run(cfg *config.Config, resetInstancesInput *service.ResetInstancesInput) error {

	qcService, err := service.Init(cfg)
	if err != nil {
		return err
	}
	instanceService, err := qcService.Instance(cfg.Zone)
	if err != nil {
		return err
	}
	instances := strings.Split(instances, ",")
	for i := range instances {
		resetInstancesInput.Instances = append(resetInstancesInput.Instances, &instances[i])
	}
	log.Debug().Interface("reset input params", resetInstancesInput).Send()
	resetInstancesOutput, err := instanceService.ResetInstances(resetInstancesInput)
	if err != nil {
		return err
	}
	if resetInstancesOutput.RetCode != nil && *resetInstancesOutput.RetCode != 0 && resetInstancesOutput.Message != nil {
		return errors.New(*resetInstancesOutput.Message)
	}
	log.Debug().Msg("reset qing yun node success")
	return nil
}
