package utils

import (
	treasurytypes "github.com/dreanity/saturn/x/treasury/types"
)

func SplitGasPriceList(gasPriceList []treasurytypes.GasPrice, paramsNumber int) [][]treasurytypes.GasPrice {
	maxBalancesPerSlice := maxPostgreSQLParams / paramsNumber
	slices := make([][]treasurytypes.GasPrice, len(gasPriceList)/maxBalancesPerSlice+1)

	sliceIndex := 0
	for index, gasPrice := range gasPriceList {
		slices[sliceIndex] = append(slices[sliceIndex], gasPrice)

		if index > 0 && index%(maxBalancesPerSlice-1) == 0 {
			sliceIndex++
		}
	}

	return slices
}

func SplitGasBidList(gasBidList []treasurytypes.GasBid, paramsNumber int) [][]treasurytypes.GasBid {
	maxBalancesPerSlice := maxPostgreSQLParams / paramsNumber
	slices := make([][]treasurytypes.GasBid, len(gasBidList)/maxBalancesPerSlice+1)

	sliceIndex := 0
	for index, gasBid := range gasBidList {
		slices[sliceIndex] = append(slices[sliceIndex], gasBid)

		if index > 0 && index%(maxBalancesPerSlice-1) == 0 {
			sliceIndex++
		}
	}

	return slices
}
