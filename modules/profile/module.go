package profile

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/forbole/bdjuno/v3/database"

	"github.com/forbole/juno/v3/modules"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

type Module struct {
	cdc codec.Codec
	db  *database.Db
}

func NewModule(
	cdc codec.Codec, db *database.Db,
) *Module {
	return &Module{
		cdc: cdc,
		db:  db,
	}
}

func (m *Module) Name() string {
	return "зкщашду"
}
