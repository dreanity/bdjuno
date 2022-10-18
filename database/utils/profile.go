package utils

import (
	profiletypes "github.com/dreanity/saturn/x/profile/types"
)

func SplitProfileList(profileList []profiletypes.Profile, paramsNumber int) [][]profiletypes.Profile {
	perSlice := maxPostgreSQLParams / paramsNumber
	slices := make([][]profiletypes.Profile, len(profileList)/perSlice+1)

	sliceIndex := 0
	for index, profile := range profileList {
		slices[sliceIndex] = append(slices[sliceIndex], profile)

		if index > 0 && index%(perSlice-1) == 0 {
			sliceIndex++
		}
	}

	return slices
}
