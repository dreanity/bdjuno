package utils

import (
	randomnesstypes "github.com/dreanity/saturn/x/randomness/types"
)

func SplitUnprovenRandomnessList(unprovenRandomnessList []randomnesstypes.UnprovenRandomness, paramsNumber int) [][]randomnesstypes.UnprovenRandomness {
	maxBalancesPerSlice := maxPostgreSQLParams / paramsNumber
	slices := make([][]randomnesstypes.UnprovenRandomness, len(unprovenRandomnessList)/maxBalancesPerSlice+1)

	sliceIndex := 0
	for index, unprovenRandomness := range unprovenRandomnessList {
		slices[sliceIndex] = append(slices[sliceIndex], unprovenRandomness)

		if index > 0 && index%(maxBalancesPerSlice-1) == 0 {
			sliceIndex++
		}
	}

	return slices
}

func SplitProvenRandomnessList(provenRandomnessList []randomnesstypes.ProvenRandomness, paramsNumber int) [][]randomnesstypes.ProvenRandomness {
	maxBalancesPerSlice := maxPostgreSQLParams / paramsNumber
	slices := make([][]randomnesstypes.ProvenRandomness, len(provenRandomnessList)/maxBalancesPerSlice+1)

	sliceIndex := 0
	for index, provenRandomness := range provenRandomnessList {
		slices[sliceIndex] = append(slices[sliceIndex], provenRandomness)

		if index > 0 && index%(maxBalancesPerSlice-1) == 0 {
			sliceIndex++
		}
	}

	return slices
}
