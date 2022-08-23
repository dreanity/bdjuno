package utils

import (
	giveawaytypes "github.com/dreanity/saturn/x/giveaway/types"
)

func SplitGiveawayList(giveawayList []giveawaytypes.Giveaway, paramsNumber int) [][]giveawaytypes.Giveaway {
	maxBalancesPerSlice := maxPostgreSQLParams / paramsNumber
	slices := make([][]giveawaytypes.Giveaway, len(giveawayList)/maxBalancesPerSlice+1)

	sliceIndex := 0
	for index, giveaway := range giveawayList {
		slices[sliceIndex] = append(slices[sliceIndex], giveaway)

		if index > 0 && index%(maxBalancesPerSlice-1) == 0 {
			sliceIndex++
		}
	}

	return slices
}
