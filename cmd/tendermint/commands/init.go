package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	//clearup for version command.
	//os.go only depends on third-party dependencies.
	tmos "github.com/YunweiMao/tendermint/libs/os"

	//only depends on libs/os/os.go
	cfg "github.com/YunweiMao/tendermint/config"

	//time.go only depends on the third-party dependencies
	tmtime "github.com/YunweiMao/tendermint/types/time"

	//random.go only depends on libs/sync/sync.go
	//sync.go only depends on third-party dependencies.
	tmrand "github.com/YunweiMao/tendermint/libs/rand"

	//privval/file.go has a very complicated dependency.
	// but you can refer more details in the file.go
	"github.com/YunweiMao/tendermint/privval"

	//p2p/genesis.go has a simple dependency
	// more details are on the file itself
	"github.com/YunweiMao/tendermint/p2p"

	//types/genesis.go is the only file we need
	"github.com/YunweiMao/tendermint/types"
)

// InitFilesCmd initialises a fresh Tendermint Core instance.
var InitFilesCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Tendermint",
	RunE:  initFiles,
}

func initFiles(cmd *cobra.Command, args []string) error {
	return initFilesWithConfig(config)
}

func initFilesWithConfig(config *cfg.Config) error {
	// private validator
	privValKeyFile := config.PrivValidatorKeyFile()
	privValStateFile := config.PrivValidatorStateFile()
	var pv *privval.FilePV
	if tmos.FileExists(privValKeyFile) {
		pv = privval.LoadFilePV(privValKeyFile, privValStateFile)
		//logger is defined in /cmd/tendermint/commands/root.go
		//since root.go and init.go are in the same package.
		// we can reuse the variable defined in root.go in init.go
		logger.Info("Found private validator", 
					"keyFile", privValKeyFile,
					"stateFile", privValStateFile)
	} else {
		pv = privval.GenFilePV(privValKeyFile, privValStateFile)
		pv.Save()
		logger.Info("Generated private validator", 
					"keyFile", privValKeyFile,
					"stateFile", privValStateFile)
	}

	//node key file
	nodeKeyFile := config.NodeKeyFile()
	if tmos.FileExists(nodeKeyFile) {
		logger.Info("Found node key", "path", nodeKeyFile)
	} else {
		if _, err := p2p.LoadOrGenNodeKey(nodeKeyFile); err != nil {
			return err
		}
		logger.Info("Generated node key", "path", nodeKeyFile)
	}

	// genesis file
	genFile := config.GenesisFile()
	if tmos.FileExists(genFile) {
		logger.Info("Found genesis file", "path", genFile)
	} else {
		genDoc := types.GenesisDoc{
			ChainID:         fmt.Sprintf("test-chain-%v", tmrand.Str(6)),
			GenesisTime:     tmtime.Now(),
			ConsensusParams: types.DefaultConsensusParams(),
		}
		pubKey, err := pv.GetPubKey()
		if err != nil {
			return fmt.Errorf("can't get pubkey: %w", err)
		}
		genDoc.Validators = []types.GenesisValidator{{
			Address: pubKey.Address(),
			PubKey:  pubKey,
			Power:   10,
		}}

		if err := genDoc.SaveAs(genFile); err != nil {
			return err
		}
		logger.Info("Generated genesis file", "path", genFile)
	}

	return nil
}
