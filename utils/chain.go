package utils

import (
	"strings"

	"github.com/forbole/bdtool/types"
)

func GetChainInfo() *types.Chain {
	chainName := GetInput("Chain Name")
	chainType := GetInput("Chain Type (mainnet, testnet, or other)")

	return &types.Chain{
		Name: strings.ToLower(chainName),
		Type: strings.ToLower(chainType),
	}
}
