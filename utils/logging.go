package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	Red    = color.New(color.FgRed).SprintFunc()
	Yellow = color.New(color.FgYellow).SprintFunc()
	Green  = color.New(color.FgGreen).SprintFunc()

	ErrorSymbol   = "❌"
	WarningSymbol = "⚠️"
	SuccessSymbol = "✔"
	ToolSymbol    = "🔧"
)

func PrintError(header string, err error) {
	fmt.Printf("%s ", Red(ErrorSymbol, " ", header))
	fmt.Printf("%s\n", err)
}

func PrintFatalError(header string, err error) {
	PrintError(header, err)
	os.Exit(-1)
}

func PrintWarning(message string) {
	fmt.Printf("%s\n", Yellow(WarningSymbol, " ", message))
}

func PrintSuccess(message string) {
	fmt.Printf("%s\n", Green(SuccessSymbol, " ", message))
}
