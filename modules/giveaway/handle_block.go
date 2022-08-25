package giveaway

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"
	"github.com/gogo/protobuf/proto"

	giveawaytypes "github.com/dreanity/saturn/x/giveaway/types"
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
		case *giveawaytypes.GiveawayCreated:
			if err := m.db.SaveGiveawayFromGiveawayCreatedEvent(_msg); err != nil {
				return err
			}
		case *giveawaytypes.GiveawayWinnersDeterminationBegun:
			if err := m.db.UpdateGiveawayFromWinnersDeterminationBegunEvent(_msg); err != nil {
				return err
			}
		case *giveawaytypes.GiveawayCancelledInsufTickets:
			if err := m.db.UpdateGiveawayFromCancelledInsufTicketsEvent(_msg); err != nil {
				return err
			}
		case *giveawaytypes.GiveawayWinnersDetermined:
			if err := m.db.UpdateGiveawayFromWinnersDeterminedEvent(_msg); err != nil {
				return err
			}
		}
	}

	return nil
}
