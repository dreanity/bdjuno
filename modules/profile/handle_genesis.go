package profile

import (
	"encoding/json"
	"fmt"

	profiletypes "github.com/dreanity/saturn/x/profile/types"
	"github.com/rs/zerolog/log"
	tmtypes "github.com/tendermint/tendermint/types"
)

func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "profile").Msg("parsing genesis")

	var genState profiletypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[profiletypes.ModuleName], &genState)
	if err != nil {
		return fmt.Errorf("error while unmarshaling profile state: %s", err)
	}

	err = m.db.SaveProfileListFromGenesis(genState.ProfileList)
	if err != nil {
		return fmt.Errorf("error while storing genesis profile list: %s", err)
	}

	return nil
}
