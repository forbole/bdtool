package types

type Themes struct {
	Default   string   `json:"default"`
	ThemeList []string `json:"themeList"`
	Light     Theme    `json:"light"`
	Dark      Theme    `json:"dark"`
}

type Theme struct {
	Background  Background  `json:"background"`
	Primary     Primary     `json:"primary"`
	Divider     string      `json:"divider"`
	Text        Text        `json:"text"`
	Fonts       Fonts       `json:"fonts"`
	PrimaryData PrimaryData `json:"primaryData"`
	Results     Results     `json:"results"`
	Tokenomics  Tokenomics  `json:"tokenomics"`
	Condition   Condition   `json:"condition"`
	Charts      Charts      `json:"charts"`
}

type Background struct {
	Default    string `json:"default"`
	Paper      string `json:"paper"`
	SurfaceOne string `json:"surfaceOne"`
	SurfaceTwo string `json:"surfaceTwo"`
}

type Primary struct {
	Main         string `json:"main"`
	ContrastText string `json:"contrastText"`
}

type Text struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
}

type Fonts struct {
	FontOne   string `json:"fontOne"`
	FontTwo   string `json:"fontTwo"`
	FontThree string `json:"fontThree"`
	FontFour  string `json:"fontFour"`
	FontFive  string `json:"fontFive"`
	Highlight string `json:"highlight"`
}

type PrimaryData struct {
	One   string `json:"one"`
	Two   string `json:"two"`
	Three string `json:"three"`
	Four  string `json:"four"`
}

type Results struct {
	Pass string `json:"pass"`
	Fail string `json:"fail"`
}

type Tokenomics struct {
	One   string `json:"one"`
	Two   string `json:"two"`
	Three string `json:"three"`
}

type Condition struct {
	Zero  string `json:"zero"`
	One   string `json:"one"`
	Two   string `json:"two"`
	Three string `json:"three"`
}

type Charts struct {
	Zero  string `json:"zero"`
	One   string `json:"one"`
	Two   string `json:"two"`
	Three string `json:"three"`
	Four  string `json:"four"`
	Five  string `json:"five"`
}

func NewEmptyTheme() *Theme {
	return &Theme{
		Background: Background{
			Default:    "#",
			Paper:      "#",
			SurfaceOne: "#",
			SurfaceTwo: "#",
		},
		Primary: Primary{
			Main:         "#",
			ContrastText: "#",
		},
		Divider: "#",
		Text: Text{
			Primary:   "#",
			Secondary: "#",
		},
		Fonts: Fonts{
			FontOne:   "#",
			FontTwo:   "#",
			FontThree: "#",
			FontFour:  "#",
			FontFive:  "#",
			Highlight: "#",
		},
		PrimaryData: PrimaryData{
			One:   "#",
			Two:   "#",
			Three: "#",
			Four:  "#",
		},
		Results: Results{
			Pass: "#",
			Fail: "#",
		},
		Tokenomics: Tokenomics{
			One:   "#",
			Two:   "#",
			Three: "#",
		},
		Condition: Condition{
			Zero:  "#",
			One:   "#",
			Two:   "#",
			Three: "#",
		},
		Charts: Charts{
			Zero:  "#",
			One:   "#",
			Two:   "#",
			Three: "#",
			Four:  "#",
			Five:  "#",
		},
	}
}
