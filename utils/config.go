package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/forbole/bdtool/types"
)

func GetConfig(chain *types.ChainInfo) *types.ChainConfig {
	if ConfigWithCLI() {
		return getConfigFromCLI(chain)
	}
	return getConfigFromFile()
}

func getConfigFromFile() *types.ChainConfig {
	path := GetInput("Enter config json file path")

	jsonConfig, err := os.Open(path)
	if err != nil {
		CheckError(fmt.Errorf("error while opening config file: %s", err))
	}
	defer jsonConfig.Close()

	bz, _ := ioutil.ReadAll(jsonConfig)

	var chainConfig types.ChainConfig
	err = json.Unmarshal(bz, &chainConfig)
	if err != nil {
		CheckError(fmt.Errorf("error while unmarshaling config file: %s", err))
	}

	return &chainConfig
}

func getConfigFromCLI(chain *types.ChainInfo) *types.ChainConfig {
	title := GetInput("Title (e.g. Desmos Block Explorer)")
	network := GetInput("Network Name (e.g. desmos-mainnet)")
	prefix := GetInput("Chain Prefix (e.g. desmos)")
	genesisTime := GetInput("Genesis Time (e.g. 2021-08-31T15:00:00)")
	genesisHeight := GetInput("Genesis Height")
	primaryTokenUnit := GetInput("Primary Token Unit (e.g. udsm)")
	votingPowerTokenUnit := GetInput("Voting Power Token Unit (e.g. udsm)")
	tokenUnits := getTokenUnits()
	graphqlEndpoint := GetInput("GraphQL Endpoint (e.g. https://example.com/v1/graphql)")
	graphqlWs := GetInput("GraphQL Web Socket (e.g. wss://example.com/v1/graphql)")
	publicRpcWs := GetInput("Public RPC Websocket (e.g. wss://rpc.example.com/websocket)")
	basePath := GetInput("URL Base Path (e.g. /desmos)")
	previewImage := GetInput("Preview Image (e.g. https://s3.example.com/chain.png)")
	matomoURL := GetInput("Matomo URL (e.g. https://example.bigdipper.live)")
	matomoSiteID := GetInput("Matomo Site ID (e.g. 1)")
	defaultTheme := GetInput("Default Theme (e.g. light)")
	themeList := getThemeList(defaultTheme)

	fmt.Printf("\x1b[%dm%s\x1b[0m", 34, "Configuring colors, no need to specify '#' before the hex code")
	lightTheme := getTheme("light")
	darkTheme := getTheme("dark")

	return &types.ChainConfig{
		Title:   title,
		Network: network,
		Icon:    fmt.Sprintf("/%s/images/%s/icon.svg", chain.Name, chain.Name),
		Logo: types.Logo{
			Default: fmt.Sprintf("/%s/images/%s/logo.svg", chain.Name, chain.Name),
		},
		Prefix: types.Prefix{
			Consensus: prefix + "valcons",
			Validator: prefix + "valoper",
			Account:   prefix,
		},
		Genesis: types.Genesis{
			Time:   genesisTime,
			Height: genesisHeight,
		},
		PrimaryTokenUnit:     primaryTokenUnit,
		VotingPowerTokenUnit: votingPowerTokenUnit,
		TokenUnits:           tokenUnits,
		Extra:                types.Extra{Profile: true},
		Endpoints: types.Endpoints{
			Graphql:            graphqlEndpoint,
			GraphqlWebsocket:   graphqlWs,
			PublicRPCWebsocket: publicRpcWs,
		},
		General: types.General{
			BasePath:     basePath,
			PreviewImage: previewImage,
		},
		Marketing: types.Marketing{
			MatomoURL:    matomoURL,
			MatomoSiteID: matomoSiteID,
		},
		Style: types.Style{
			Themes: types.Themes{
				Default:   defaultTheme,
				ThemeList: themeList,
				Light:     lightTheme,
				Dark:      darkTheme,
			},
		},
	}

}
func GetConfigBz(config *types.ChainConfig) []byte {
	bz, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		CheckError(fmt.Errorf("error while making json: %s", err))
	}

	return bz
}
