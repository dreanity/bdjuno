package giveaway

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"

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

		default:
			_msg.ProtoMessage()
		}
	}

	return nil
}
