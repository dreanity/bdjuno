package treasury

import (
	"encoding/json"
	"fmt"

	treasurytypes "github.com/dreanity/saturn/x/treasury/types"
	"github.com/rs/zerolog/log"
	tmtypes "github.com/tendermint/tendermint/types"
)

func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "treasury").Msg("parsing genesis")

	var genState treasurytypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[treasurytypes.ModuleName], &genState)
	if err != nil {
		return fmt.Errorf("error while unmarshaling treasury state: %s", err)
	}

	err = m.db.SaveTreasurer(genState.Treasurer)
	if err != nil {
		return fmt.Errorf("error while storing genesis treasurer: %s", err)
	}

	err = m.db.SaveGasPriceListFromGenesis(genState.GasPriceList)
	if err != nil {
		return fmt.Errorf("error while storing genesis gas price list: %s", err)
	}

	err = m.db.SaveGasBidListFromGenesis(genState.GasBidList)
	if err != nil {
		return fmt.Errorf("error while storing genesis gas bid list: %s", err)
	}

	return nil
}
