package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/forbole/bdtool/types"
	"github.com/manifoldco/promptui"
)

func GetInput(question string) string {
	prompt := promptui.Prompt{
		Label: question,
	}

	result, err := prompt.Run()
	if err != nil {
		CheckError(fmt.Errorf("prompt failed %v", err))
	}

	return result
}

func GetTokenInput(question string) string {
	instruction := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, "Enter GitHub Personal Access Token. You can generate one at https://github.com/settings/tokens/new")
	fmt.Println(instruction)

	prompt := promptui.Prompt{
		Label: question,
		Mask:  '*',
	}

	result, err := prompt.Run()
	if err != nil {
		CheckError(fmt.Errorf("prompt failed %v", err))
	}

	return result
}

func getColorInput(question string) string {
	colorCode := GetInput(question)
	return "#" + colorCode
}

func GetBool(question string) bool {
	prompt := promptui.Select{
		Label: question,
		Items: []string{"YES", "NO"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		CheckError(fmt.Errorf("prompt failed %v", err))
	}

	if result == "YES" {
		return true
	}

	return false
}

func ConfigWithCLI() bool {
	prompt := promptui.Select{
		Label: "Choose a method to create config file",
		Items: []string{"config with CLI inputs", "import existing config json file"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		CheckError(fmt.Errorf("prompt failed %v", err))
	}

	if result == "config with CLI inputs" {
		return true
	}

	return false
}

func getTokenUnits() map[string]types.Token {
	var tokenUnits = map[string]types.Token{}
	i := 1
	for {
		tokenUnit := GetInput(fmt.Sprintf("Token Unit %v (udsm)", i))
		display := GetInput("Display (dsm)")
		exponent := GetInput("Exponent (6)")
		exponentInt, err := strconv.ParseInt(exponent, 10, 16)
		if err != nil {
			CheckError(err)
		}

		tokenUnits[tokenUnit] = types.Token{
			Display:  display,
			Exponent: int16(exponentInt),
		}

		moreTokenUnits := GetBool("More token units")
		if !moreTokenUnits {
			break
		}

		i++
	}

	return tokenUnits

}

func getThemeList(defaultTheme string) []string {
	var themeList = []string{defaultTheme}

	hasDeuteranopia := GetBool("Apply deuteranopia theme")
	if hasDeuteranopia {
		themeList = append(themeList, "deuteranopia")
	}

	hasTritanopia := GetBool("Apply tritanopia theme")
	if hasTritanopia {
		themeList = append(themeList, "tritanopia")
	}

	return themeList
}

func getTheme(themeType string) types.Theme {
	hasTheme := GetBool(fmt.Sprintf("Configure %s theme", themeType))
	if !hasTheme {
		return *types.NewEmptyTheme()
	}
	// Background
	defaultBackground := getColorInput("Default Background")
	paper := getColorInput("Paper")
	surfaceOne := getColorInput("Surface One")
	surfaceTwo := getColorInput("Surface Two")

	// Primary
	main := getColorInput("Main")
	contrastText := getColorInput("Contrast Text")

	divider := getColorInput("divider")

	// Text
	primary := getColorInput("primary")
	secondary := getColorInput("secondary")

	// Fonts
	fontOne := getColorInput("Font One")
	fontTwo := getColorInput("Font Two")
	fontThree := getColorInput("Font Three")
	fontFour := getColorInput("Font Four")
	fontFive := getColorInput("Font Five")
	highlight := getColorInput("Highlight")

	// PrimaryData
	primaryDataOne := getColorInput("Primary Data One")
	primaryDataTwo := getColorInput("Primary Data Two")
	primaryDataThree := getColorInput("Primary Data Three")
	primaryDataFour := getColorInput("Primary Data Four")

	// Results
	resultPass := getColorInput("Result Pass")
	resultFail := getColorInput("Result Fail")

	// Tokenomics
	tokenomicsOne := getColorInput("Tokenomics One")
	tokenomicsTwo := getColorInput("Tokenomics Two")
	tokenomicsThree := getColorInput("Tokenomics Three")

	// Condition
	conditionZero := getColorInput("Condition Zero")
	conditionOne := getColorInput("Condition One")
	conditionTwo := getColorInput("Condition Two")
	conditionThree := getColorInput("Condition Three")

	// Charts
	chartsZero := getColorInput("Charts Zero")
	chartsOne := getColorInput("Charts One")
	chartsTwo := getColorInput("Charts Two")
	chartsThree := getColorInput("Charts Three")
	chartsFour := getColorInput("Charts Four")
	chartsFive := getColorInput("Charts Five")

	return types.Theme{
		Background: types.Background{
			Default:    defaultBackground,
			Paper:      paper,
			SurfaceOne: surfaceOne,
			SurfaceTwo: surfaceTwo,
		},
		Primary: types.Primary{
			Main:         main,
			ContrastText: contrastText,
		},
		Divider: divider,
		Text: types.Text{
			Primary:   primary,
			Secondary: secondary,
		},
		Fonts: types.Fonts{
			FontOne:   fontOne,
			FontTwo:   fontTwo,
			FontThree: fontThree,
			FontFour:  fontFour,
			FontFive:  fontFive,
			Highlight: highlight,
		},
		PrimaryData: types.PrimaryData{
			One:   primaryDataOne,
			Two:   primaryDataTwo,
			Three: primaryDataThree,
			Four:  primaryDataFour,
		},
		Results: types.Results{
			Pass: resultPass,
			Fail: resultFail,
		},
		Tokenomics: types.Tokenomics{
			One:   tokenomicsOne,
			Two:   tokenomicsTwo,
			Three: tokenomicsThree,
		},
		Condition: types.Condition{
			Zero:  conditionZero,
			One:   conditionOne,
			Two:   conditionTwo,
			Three: conditionThree,
		},
		Charts: types.Charts{
			Zero:  chartsZero,
			One:   chartsOne,
			Two:   chartsTwo,
			Three: chartsThree,
			Four:  chartsFour,
			Five:  chartsFive,
		},
	}
}

// CheckError should be used to naively panics if an error is not nil.
func CheckError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
