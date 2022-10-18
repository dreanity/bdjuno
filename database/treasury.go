package database

import (
	"fmt"

	treasurytypes "github.com/dreanity/saturn/x/treasury/types"
	dbutils "github.com/forbole/bdjuno/v3/database/utils"
)

func (db *Db) SaveTreasurer(treasurer treasurytypes.Treasurer) error {
	stmt := `
INSERT INTO treasurer (address) 
VALUES ($1)
ON CONFLICT (address) DO UPDATE 
    SET address = excluded.address
WHERE treasurer.address = excluded.address`

	_, err := db.Sql.Exec(stmt, treasurer.Address)
	if err != nil {
		return fmt.Errorf("error while storing treasurer: %s", err)
	}

	return nil
}

// ---------------------------------------------------------------------------

func (db *Db) SaveGasPriceListFromGenesis(gasPriceList []treasurytypes.GasPrice) error {
	paramsNumber := 1
	slices := dbutils.SplitGasPriceList(gasPriceList, paramsNumber)

	for _, list := range slices {
		if len(list) == 0 {
			continue
		}

		// Store up-to-date data
		err := db.saveGasPriceList(paramsNumber, list)
		if err != nil {
			return fmt.Errorf("error while storing gas price list: %s", err)
		}
	}

	return nil
}

func (db *Db) saveGasPriceList(paramsNumber int, gasPriceList []treasurytypes.GasPrice) error {
	if len(gasPriceList) == 0 {
		return nil
	}

	stmt := `INSERT INTO gas_price (chain, token_address, token_symbol, value) VALUES `
	var params []interface{}

	for i, gasPrice := range gasPriceList {
		ai := i * paramsNumber
		stmt += fmt.Sprintf("($%d, $%d, $%d, $%d),", ai+1, ai+2, ai+3, ai+4)
		params = append(params, gasPrice.Chain, gasPrice.TokenAddress, gasPrice.TokenSymbol, gasPrice.Value)
	}

	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`
	_, err := db.Sql.Exec(stmt, params...)
	if err != nil {
		return fmt.Errorf("error while storing gas price list: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------

func (db *Db) SaveGasBidListFromGenesis(gasBidList []treasurytypes.GasBid) error {
	paramsNumber := 1
	slices := dbutils.SplitGasBidList(gasBidList, paramsNumber)

	for _, list := range slices {
		if len(list) == 0 {
			continue
		}

		// Store up-to-date data
		err := db.saveGasBidList(paramsNumber, list)
		if err != nil {
			return fmt.Errorf("error while storing gas bid list: %s", err)
		}
	}

	return nil
}

func (db *Db) saveGasBidList(paramsNumber int, gasBidList []treasurytypes.GasBid) error {
	if len(gasBidList) == 0 {
		return nil
	}

	stmt := `INSERT INTO gas_bid (chain, number) VALUES `
	var params []interface{}

	for i, gasBid := range gasBidList {
		ai := i * paramsNumber
		stmt += fmt.Sprintf("($%d, $%d),", ai+1, ai+2)
		params = append(params, gasBid.Chain, gasBid.Number)
	}

	stmt = stmt[:len(stmt)-1]
	stmt += ` ON CONFLICT DO NOTHING`
	_, err := db.Sql.Exec(stmt, params...)
	if err != nil {
		return fmt.Errorf("error while storing gas bid list: %s", err)
	}

	return nil
}

// ----------------------------------------------------------

func (db *Db) UpdateGasPriceFromGasPricesChangedEvent(event *treasurytypes.GasPricesChanged) error {
	if len(event.GasPrices) == 0 {
		return nil
	}

	stmt := `UPDATE gas_price 
    SET chain = $1,
		token_address = $2,
		token_symbol = $3,
		value = $4,
WHERE gas_price.token_address = $2`
	var params []interface{}

	for i, gasPrice := range event.GasPrices {
		stmt += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d),", i+1, i+2, i+3, i+4, i+5)
		params = append(params, gasPrice.Chain, gasPrice.TokenAddress, gasPrice.TokenSymbol, gasPrice.Value)
	}

	stmt = stmt[:len(stmt)-1]
	_, err := db.Sql.Exec(stmt, params...)
	if err != nil {
		return fmt.Errorf("error while updating gas price from gas prices changed event: %s", err)
	}

	return nil
}

// -----------------------------------------------

func (db *Db) SetGasBidFromGasBidExecutedEvent(event *treasurytypes.GasBidExecuted) error {
	stmt := `INSERT INTO gas_bid (
		chain,
		number
	) VALUES ($1, $2)
ON CONFLICT (chain) DO UPDATE 
    SET chain = excluded.chain,
        number = excluded.number,
WHERE gas_bid.chain = excluded.chain`

	_, err := db.Sql.Exec(stmt,
		event.Chain,
		event.BidNumber,
	)
	if err != nil {
		return fmt.Errorf("error while storing giveaway from giveaway created event: %s", err)
	}

	return nil
}
