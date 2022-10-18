package randomness

import (
	"encoding/json"
	"fmt"

	randomnesstypes "github.com/dreanity/saturn/x/randomness/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/rs/zerolog/log"
)

func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "randomness").Msg("parsing genesis")

	// Read the genesis state
	var genState randomnesstypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[randomnesstypes.ModuleName], &genState)
	if err != nil {
		return fmt.Errorf("error while unmarshaling randomness state: %s", err)
	}

	err = m.db.SaveRandomnessChainInfo(genState.ChainInfo)
	if err != nil {
		return fmt.Errorf("error while storing genesis randomness chain info: %s", err)
	}

	err = m.db.SaveUnprovenRandomnessListFromGenesis(genState.UnprovenRandomnessList)
	if err != nil {
		return fmt.Errorf("error while storing genesis unproven randomness list: %s", err)
	}

	err = m.db.SaveProvenRandomnessListFromGenesis(genState.ProvenRandomnessList)
	if err != nil {
		return fmt.Errorf("error while storing genesis proven randomness list: %s", err)
	}

	return nil
}
