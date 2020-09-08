package main

import (
	"fmt"
	"os"
	"strings"
)

type locale struct {
	language string
	region   string
}

func getSupportedLocales() map[locale]bool {
	supportedLocales := make(map[locale]bool, 5)
	supportedLocales[locale{"en", "US"}] = true
	supportedLocales[locale{"en", "CA"}] = true
	supportedLocales[locale{"fr", "CA"}] = true
	supportedLocales[locale{"fr", "FR"}] = true
	supportedLocales[locale{"ru", "RU"}] = true
	return supportedLocales
}

func (locale locale) supported() bool {
	supportedLocales := getSupportedLocales()
	return supportedLocales[locale]
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Must provide exactly one argument: the locale.")
		os.Exit(1)
	}
	localeArg := os.Args[1]
	localeParts := strings.Split(localeArg, "_")
	if len(localeParts) != 2 {
		fmt.Println("Error: The locale must have exactly two parts: the language and region, separated by a single underscore.")
		os.Exit(1)
	}
	locale := locale{
		language: localeParts[0],
		region:   localeParts[1],
	}
	if locale.supported() {
		fmt.Printf("Locale %v is supported.\n", localeArg)
		os.Exit(0)
	} else {
		fmt.Printf("Locale %v is not supported.\n", localeArg)
		os.Exit(1)
	}
}
