package options

import (
	"errors"
	"os"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
	"github.com/yunify/qingcloud-sdk-go/service"
)

func MakeAttachKeyPairs() *cobra.Command {
	command := &cobra.Command{
		Use:     "attach",
		Example: `qingyun attach -h`,
	}
	if err != nil {
		log.Error().Err(err).Send()
		return command
	}
	flags(command)
	command.Flags().String("instances", "", "一个或多个主机ID,eg:i-66655,i-a44w5")
	command.Flags().String("key-pairs", "", "eg:is_anbe1kw5,if_anbe1kw5")
	command.Run = func(cmd *cobra.Command, _ []string) {
		assignment(cfg)
		if err := runAttachKeyPairs(cmd); err != nil {
			log.Error().Err(err).Msg("attach ssh key pairs failed")
			os.Exit(1)
		}
	}
	return command
}

func runAttachKeyPairs(cmd *cobra.Command) error {
	var (
		instancesSlice, keyPairSlice []*string
	)

	ins, _ := cmd.Flags().GetString("instances")
	kps, _ := cmd.Flags().GetString("key-pairs")

	instances := strings.Split(ins, ",")
	for i := range instances {
		instancesSlice = append(instancesSlice, &instances[i])
	}

	keyPairs := strings.Split(kps, ",")
	for i := range keyPairs {
		keyPairSlice = append(keyPairSlice, &keyPairs[i])
	}
	detachKeyPairsInput := service.AttachKeyPairsInput{
		Instances: instancesSlice,
		KeyPairs:  keyPairSlice,
	}

	qingCloudService, err := service.Init(cfg)
	if err != nil {
		return err
	}
	keyPairService, err := qingCloudService.KeyPair(cfg.Zone)
	if err != nil {
		return err
	}

	attachKeyPairsOutput, err := keyPairService.AttachKeyPairs(&detachKeyPairsInput)
	if err != nil {
		return err
	}
	log.Info().Interface("attachKeyPairsOutput", attachKeyPairsOutput).Send()
	if attachKeyPairsOutput.RetCode != nil && *attachKeyPairsOutput.RetCode != 0 && attachKeyPairsOutput.Message != nil {
		return errors.New(*attachKeyPairsOutput.Message)
	}
	log.Debug().Msg("attach ssh key pairs success")
	return nil
}
