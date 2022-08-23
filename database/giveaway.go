package database

import (
	"encoding/json"
	"fmt"

	giveawaytypes "github.com/dreanity/saturn/x/giveaway/types"
	dbutils "github.com/forbole/bdjuno/v3/database/utils"
	"github.com/lib/pq"
)

func (db *Db) SaveGiveawayList(giveawayList []giveawaytypes.Giveaway, ticketCountList []giveawaytypes.TicketCount) error {
	paramsNumber := 9
	slices := dbutils.SplitGiveawayList(giveawayList, paramsNumber)

	for _, list := range slices {
		if len(list) == 0 {
			continue
		}

		// Store up-to-date data
		err := db.saveGiveawayList(paramsNumber, list, ticketCountList)
		if err != nil {
			return fmt.Errorf("error while storing unproven randomness list: %s", err)
		}
	}

	return nil
}

func (db *Db) saveGiveawayList(paramsNumber int, giveawayList []giveawaytypes.Giveaway, ticketCountList []giveawaytypes.TicketCount) error {
	if len(giveawayList) == 0 {
		return nil
	}

	stmt := `INSERT INTO giveaway (
		index, 
		duration, 
		created_at, 
		name, 
		completion_height, 
		winning_ticket_numbers, 
		prizes, 
		status,
		ticket_count
	) VALUES `

	var params []interface{}
	for i, giveaway := range giveawayList {
		ai := i * paramsNumber
		stmt += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d),",
			ai+1, ai+2, ai+3, ai+4, ai+5, ai+6, ai+7, ai+8, ai+9)

		winningTicketNumbersInt := make([]int32, 0)

		for _, winningTicketNumber := range giveaway.WinningTicketNumbers {
			winningTicketNumbersInt = append(winningTicketNumbersInt, int32(winningTicketNumber))
		}

		prizesJSONBytes, err := json.Marshal(giveaway.Prizes)
		if err != nil {
			return fmt.Errorf("error while marshaling giveaway prizes: %s", err)
		}

		ticketCount := 0
		for _, tikcetCountGiveaway := range ticketCountList {
			if tikcetCountGiveaway.GiveawayId == giveaway.Index {
				ticketCount = int(tikcetCountGiveaway.Count)
			}
		}

		params = append(
			params,
			giveaway.Index,
			giveaway.Duration,
			giveaway.CreatedAt,
			giveaway.Name,
			giveaway.CompletionHeight,
			pq.Int32Array(winningTicketNumbersInt),
			string(prizesJSONBytes),
			giveaway.Status,
			ticketCount,
		)
	}

	stmt = stmt[:len(stmt)-1]
	stmt += " ON CONFLICT DO NOTHING"
	_, err := db.Sql.Exec(stmt, params...)
	if err != nil {
		return fmt.Errorf("error while storing giveaway list: %s", err)
	}

	return nil
}

// ----------------------------------------------------

func (db *Db) SaveGiveawayFromEvent(event *giveawaytypes.GiveawayCreated) error {
	stmt := `INSERT INTO giveaway (
		index, 
		duration, 
		created_at, 
		name, 
		completion_height, 
		winning_ticket_numbers, 
		prizes, 
		status,
		ticket_count
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT (round) DO UPDATE 
    SET index = excluded.index,
        duration = excluded.duration,
		created_at = excluded.created_at,
		name = excluded.name,
		completion_height = excluded.completion_height,
		winning_ticket_numbers = excluded.winning_ticket_numbers,
		prizes = excluded.prizes,
		status = excluded.status,
		ticket_count = excluded.ticket_count
WHERE giveaway.index = excluded.index`

	winningTicketNumbersInt := make([]int32, 0)

	for _, winningTicketNumber := range event.WinningTicketNumbers {
		winningTicketNumbersInt = append(winningTicketNumbersInt, int32(winningTicketNumber))
	}

	prizesJSONBytes, err := json.Marshal(event.Prizes)
	if err != nil {
		return fmt.Errorf("error while marshaling giveaway prizes: %s", err)
	}

	_, err = db.Sql.Exec(stmt,
		event.Index,
		event.Duration,
		event.CreatedAt,
		event.Name,
		event.CompletionHeight,
		pq.Int32Array(winningTicketNumbersInt),
		string(prizesJSONBytes),
		event.Status,
		0,
	)
	if err != nil {
		return fmt.Errorf("error while storing giveaway from event: %s", err)
	}

	return nil
}

// -----------------------------------------------------------------------------

// func (db *Db) ChangeFromEvent(event *giveawaytypes.GiveawayCreated) error {}
