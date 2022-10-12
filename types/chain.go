package types

type ChainInfo struct {
	Name string
	Type string
}

type ChainConfig struct {
	Title                string           `json:"title"`
	Network              string           `json:"network"`
	Icon                 string           `json:"icon"`
	Logo                 Logo             `json:"logo"`
	Prefix               Prefix           `json:"prefix"`
	Genesis              Genesis          `json:"genesis"`
	PrimaryTokenUnit     string           `json:"primaryTokenUnit"`
	VotingPowerTokenUnit string           `json:"votingPowerTokenUnit"`
	TokenUnits           map[string]Token `json:"tokenUnits"`
	Extra                Extra            `json:"extra"`
	Endpoints            Endpoints        `json:"endpoints"`
	General              General          `json:"general"`
	Marketing            Marketing        `json:"marketing"`
	Style                Style            `json:"style"`
}

type Logo struct {
	Default string `json:"default"`
}

type Prefix struct {
	Consensus string `json:"consensus"`
	Validator string `json:"validator"`
	Account   string `json:"account"`
}

type Genesis struct {
	Time   string      `json:"time"`
	Height interface{} `json:"height"`
}

type Token struct {
	Display  string `json:"display"`
	Exponent int16  `json:"exponent"`
}

type Extra struct {
	Profile bool `json:"profile"`
}

type Endpoints struct {
	Graphql            string `json:"graphql"`
	GraphqlWebsocket   string `json:"graphqlWebsocket"`
	PublicRpcWebsocket string `json:"publicRpcWebsocket"`
}

type General struct {
	BasePath     string `json:"basePath"`
	PreviewImage string `json:"previewImage"`
}

type Marketing struct {
	MatomoURL    string `json:"matomoURL"`
	MatomoSiteID string `json:"matomoSiteID"`
}

type Style struct {
	Themes Themes `json:"themes"`
}
