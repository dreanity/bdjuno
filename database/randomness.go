package database

import (
	"fmt"

	randomnesstypes "github.com/dreanity/saturn/x/randomness/types"
	dbutils "github.com/forbole/bdjuno/v3/database/utils"
)

const (
	maxPostgreSQLParams = 65535
)

func (db *Db) SaveRandomnessChainInfo(chainInfo randomnesstypes.ChainInfo) error {
	stmt := `
INSERT INTO randomness_chain_info (public_key, period, genesis_time, hash) 
VALUES ($1, $2, $3, $4)
ON CONFLICT (hash) DO UPDATE 
    SET public_key = excluded.public_key,
        period = excluded.period,
		genesis_time = excluded.genesis_time,
		hash = excluded.hash
WHERE randomness_chain_info.hash = excluded.hash`

	_, err := db.Sql.Exec(stmt, chainInfo.PublicKey, chainInfo.Period, chainInfo.GenesisTime, chainInfo.Hash)
	if err != nil {
		return fmt.Errorf("error while storing randomness chain info: %s", err)
	}

	return nil
}

// -----------------------------------------------------------

func (db *Db) SaveUnprovenRandomnessList(unprovenRandomnessList []randomnesstypes.UnprovenRandomness) error {
	paramsNumber := 1
	slices := dbutils.SplitUnprovenRandomnessList(unprovenRandomnessList, paramsNumber)

	for _, list := range slices {
		if len(list) == 0 {
			continue
		}

		// Store up-to-date data
		err := db.saveUnprovenRandomnessList(paramsNumber, list)
		if err != nil {
			return fmt.Errorf("error while storing unproven randomness list: %s", err)
		}
	}

	return nil
}

func (db *Db) SaveUnprovenRandomnessFromEvent(event *randomnesstypes.UnprovenRandomnessCreated) error {
	stmt := `INSERT INTO unproven_randomness (round) VALUES ($1) ON CONFLICT DO NOTHING`

	_, err := db.Sql.Exec(stmt, event.Round)
	if err != nil {
		return fmt.Errorf("error while storing unproven randomness from event: %s", err)
	}

	return nil
}

func (db *Db) saveUnprovenRandomnessList(paramsNumber int, unprovenRandomnessList []randomnesstypes.UnprovenRandomness) error {
	if len(unprovenRandomnessList) == 0 {
		return nil
	}

	stmt := `INSERT INTO unproven_randomness (round) VALUES `
	var params []interface{}

	for i, unprovenRandomness := range unprovenRandomnessList {
		ai := i * paramsNumber
		stmt += fmt.Sprintf("($%d),", ai+1)
		params = append(params, unprovenRandomness.Round)
	}

	stmt = stmt[:len(stmt)-1]
	stmt += " ON CONFLICT DO NOTHING"
	_, err := db.Sql.Exec(stmt, params...)
	if err != nil {
		return fmt.Errorf("error while storing unproven randomness list: %s", err)
	}

	return nil
}

// -----------------------------------------------------------

func (db *Db) SaveProvenRandomnessList(provenRandomnessList []randomnesstypes.ProvenRandomness) error {
	paramsNumber := 4
	slices := dbutils.SplitProvenRandomnessList(provenRandomnessList, paramsNumber)

	for _, list := range slices {
		if len(list) == 0 {
			continue
		}

		// Store up-to-date data
		err := db.saveProvenRandomnessList(paramsNumber, list)
		if err != nil {
			return fmt.Errorf("error while storing proven randomness list: %s", err)
		}
	}

	return nil
}

func (db *Db) saveProvenRandomnessList(paramsNumber int, provenRandomnessList []randomnesstypes.ProvenRandomness) error {
	if len(provenRandomnessList) == 0 {
		return nil
	}

	stmt := `INSERT INTO proven_randomness (round, randomness, signature, previous_signature) VALUES `
	var params []interface{}

	for i, provenRandomness := range provenRandomnessList {
		ai := i * paramsNumber
		stmt += fmt.Sprintf("($%d, $%d, $%d, $%d),", ai+1, ai+2, ai+3, ai+4)
		params = append(params, provenRandomness.Round, provenRandomness.Randomness, provenRandomness.Signature, provenRandomness.PreviousSignature)
	}

	stmt = stmt[:len(stmt)-1]
	stmt += " ON CONFLICT DO NOTHING"
	_, err := db.Sql.Exec(stmt, params...)
	if err != nil {
		return fmt.Errorf("error while storing proven randomness list: %s", err)
	}

	return nil
}

func (db *Db) SaveProvenRandomnessFromEvent(event *randomnesstypes.ProvenRandomnessCreated) error {
	stmt := `INSERT INTO proven_randomness (round, randomness, signature, previous_signature) VALUES ($1, $2, $3, $4) 
ON CONFLICT (hash) DO UPDATE 
    SET round = excluded.round,
        randomness = excluded.randomness,
		signature = excluded.signature,
		previous_signature = excluded.previous_signature
WHERE proven_randomness.round = excluded.round`

	_, err := db.Sql.Exec(stmt, event.Round, event.Randomness, event.Signature, event.PreviousSignature)
	if err != nil {
		return fmt.Errorf("error while storing proven randomness from event: %s", err)
	}

	return nil
}

// -----------------------------------------------------------
