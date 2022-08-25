package giveaway

import (
	"encoding/json"
	"fmt"

	giveawaytypes "github.com/dreanity/saturn/x/giveaway/types"
	"github.com/rs/zerolog/log"
	tmtypes "github.com/tendermint/tendermint/types"
)

func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "giveaway").Msg("parsing genesis")

	var genState giveawaytypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[giveawaytypes.ModuleName], &genState)
	if err != nil {
		return fmt.Errorf("error while unmarshaling staking state: %s", err)
	}

	err = m.db.SaveGiveawayListFromGenesis(genState.GiveawayList, genState.TicketCountList)
	if err != nil {
		return fmt.Errorf("error while storing genesis giveaway list: %s", err)
	}

	err = m.db.SaveTicketListFromGenesis(genState.TicketList)
	if err != nil {
		return fmt.Errorf("error while storing genesis ticket list: %s", err)
	}

	return nil
}
