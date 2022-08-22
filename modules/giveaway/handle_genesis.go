package giveaway

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
	tmtypes "github.com/tendermint/tendermint/types"
)

func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "randomness").Msg("parsing genesis")

	return nil
}
