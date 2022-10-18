package profile

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"
	"github.com/gogo/protobuf/proto"

	profiletypes "github.com/dreanity/saturn/x/profile/types"
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
		case *profiletypes.ProfileUpdated:
			if err := m.db.SaveProfileFromProfileUpdatedEvent(_msg); err != nil {
				return err
			}
		default:
			continue
		}
	}

	return nil
}
