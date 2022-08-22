package randomness

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"

	// abci "github.com/tendermint/tendermint/abci/types"
	randomnesstypes "github.com/dreanity/saturn/x/randomness/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// HandleBlock implements BlockModule
func (m *Module) HandleBlock(
	block *tmctypes.ResultBlock, res *tmctypes.ResultBlockResults, txs []*juno.Tx, vals *tmctypes.ResultValidators,
) error {
	// events := new([]abci.Event)

	for _, event := range res.BeginBlockEvents {
		msg, err := sdk.ParseTypedEvent(event)
		if err != nil {
			return err
		}

		switch _msg := msg.(type) {

		case *randomnesstypes.UnprovenRandomnessCreated:
			if err := m.db.SaveUnprovenRandomnessFromEvent(_msg); err != nil {
				return err
			}

		case *randomnesstypes.ProvenRandomnessCreated:
			if err := m.db.SaveProvenRandomnessFromEvent(_msg); err != nil {
				return err
			}
		}
	}

	return nil
}
