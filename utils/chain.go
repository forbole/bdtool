package utils

import (
	"strings"

	"github.com/forbole/bdtool/types"
)

func GetChainInfo() *types.ChainInfo {
	chainName := GetInput("Chain Name")
	chainType := GetInput("Chain Type (mainnet, testnet, or other)")

	return &types.ChainInfo{
		Name: strings.ToLower(chainName),
		Type: strings.ToLower(chainType),
	}
}
