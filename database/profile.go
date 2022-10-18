package database

import (
	"fmt"

	profiletypes "github.com/dreanity/saturn/x/profile/types"
	dbutils "github.com/forbole/bdjuno/v3/database/utils"
)

func (db *Db) SaveProfileListFromGenesis(profileList []profiletypes.Profile) error {
	paramsNumber := 10
	slices := dbutils.SplitProfileList(profileList, paramsNumber)

	for _, list := range slices {
		if len(list) == 0 {
			continue
		}

		// Store up-to-date data
		err := db.saveProfileList(paramsNumber, list)
		if err != nil {
			return fmt.Errorf("error while storing unproven randomness list: %s", err)
		}
	}

	return nil
}

func (db *Db) saveProfileList(paramsNumber int, profileList []profiletypes.Profile) error {
	if len(profileList) == 0 {
		return nil
	}

	stmt := `INSERT INTO profile (
		address, 
		name, 
		avatar_url, 
		banner_url
	) VALUES `

	var params []interface{}
	for i, profile := range profileList {
		ai := i * paramsNumber
		stmt += fmt.Sprintf("($%d, $%d, $%d, $%d),",
			ai+1, ai+2, ai+3, ai+4)

		params = append(
			params,
			profile.Address,
			profile.Name,
			profile.AvatarUrl,
			profile.BannerUrl,
		)
	}

	stmt = stmt[:len(stmt)-1]
	stmt += " ON CONFLICT DO NOTHING"
	_, err := db.Sql.Exec(stmt, params...)
	if err != nil {
		return fmt.Errorf("error while storing profile list: %s", err)
	}

	return nil
}

// --------------------------------------------------------------------

func (db *Db) SaveProfileFromProfileUpdatedEvent(event *profiletypes.ProfileUpdated) error {
	stmt := `INSERT INTO profile (
		address, 
		name, 
		avatar_url, 
		banner_url
	) VALUES ($1, $2, $3, $4)
ON CONFLICT (address) DO UPDATE 
    SET address = excluded.address,
        name = excluded.name,
		avatar_url = excluded.avatar_url,
		banner_url = excluded.banner_url
WHERE profile.address = excluded.address`

	_, err := db.Sql.Exec(stmt,
		event.Address,
		event.Name,
		event.AvatarUrl,
		event.BannerUrl,
	)
	if err != nil {
		return fmt.Errorf("error while storing profile from profile updated event: %s", err)
	}

	return nil
}
