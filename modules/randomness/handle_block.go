package randomness

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"
	"github.com/gogo/protobuf/proto"

	randomnesstypes "github.com/dreanity/saturn/x/randomness/types"
	abci "github.com/tendermint/tendermint/abci/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// HandleBlock implements BlockModule
func (m *Module) HandleBlock(
	block *tmctypes.ResultBlock, res *tmctypes.ResultBlockResults, txs []*juno.Tx, vals *tmctypes.ResultValidators,
) error {
	var events []abci.Event

	events = append(events, res.BeginBlockEvents...)

	for _, tx := range txs {
		events = append(events, tx.Events...)
	}

	for _, event := range events {
		concreteGoType := proto.MessageType(event.Type)
		if concreteGoType == nil {
			continue
		}

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
