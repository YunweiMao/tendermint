package commands

import (
	"fmt"
	"os"
	//"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	//config.go only contain standard setup
	// config.go does not use any package in tendermint.
	cfg "github.com/YunweiMao/tendermint/config"

	//log majorly use github.com/go-kit/log thing
	"github.com/YunweiMao/tendermint/libs/log"

	//we only need setup.go file
	"github.com/YunweiMao/tendermint/libs/cli"

	tmflags "github.com/YunweiMao/tendermint/libs/cli/flags"
)

var (
	//DefaultConfig is in file /config
	config = cfg.DefaultConfig()
	//NewTMLogger is in file /libs/log/tm_logger.go
	//NewSyncWriter is in file /libs/log/logger.go
	logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
)

/*define init func
In order to define init, we need define registerFlagsRootCmd and RootCmd.
So the code following init() are defining registerFlagsRootCmd and RootCmd.
However, in order to define RootCmd, we need define a func called ParseConfig.
*/

func init() {
	registerFlagsRootCmd(RootCmd)
}

func registerFlagsRootCmd(cmd *cobra.Command) {
	cmd.PersistentFlags().String("log_level", config.LogLevel, "log level")
}

// RootCmd is the root command for Tendermint core.
var RootCmd = &cobra.Command{
	Use:   "tendermint",
	Short: "BFT state machine replication for applications in any programming languages",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if cmd.Name() == VersionCmd.Name() {
			return nil
		}

		config, err = ParseConfig(cmd)
		if err != nil {
			return err
		}

		if config.LogFormat == cfg.LogFormatJSON {
			//NewTMJSONLogger is in /libs/log/tm_json_logger.go
			logger = log.NewTMJSONLogger(log.NewSyncWriter(os.Stdout))
		}

		logger, err = tmflags.ParseLogLevel(config.LogLevel, logger, cfg.DefaultLogLevel)
		if err != nil {
			return err
		}

		if viper.GetBool(cli.TraceFlag) {
			//NewTracingLogger is in /libs/log/tracing_logger.go
			logger = log.NewTracingLogger(logger)
		}

		logger = logger.With("module", "main")
		return nil
	},
}

// ParseConfig retrieves the default environment configuration,
// sets up the Tendermint root and ensures that the root exists
func ParseConfig(cmd *cobra.Command) (*cfg.Config, error) {
	conf := cfg.DefaultConfig()
	err := viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	var home string
	if os.Getenv("TMHOME") != "" {
		home = os.Getenv("TMHOME")
	} else {
		home, err = cmd.Flags().GetString(cli.HomeFlag)
		if err != nil {
			return nil, err
		}
	}

	conf.RootDir = home

	conf.SetRoot(conf.RootDir)
	cfg.EnsureRoot(conf.RootDir)
	if err := conf.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("error in config file: %v", err)
	}
	return conf, nil
}
