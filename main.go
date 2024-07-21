package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	ESCAPTE   string = "e"
	UNESCAPTE string = "u"
)

func main() {

	var jsonPayload string
	flag.StringVar(&jsonPayload, "p", "", "Json payload")

	var formatOption string
	flag.StringVar(&formatOption, "f", "", "Formatting option")

	flag.Parse()

	if jsonPayload == "" {
		fmt.Println("Json payload (-p) is requied but not supplied.")
		os.Exit(1)
	}

	if formatOption == "" {
		fmt.Println("Format option (-f) is requied but not supplied.")
		os.Exit(1)
	}

	fmt.Println()

	if formatOption == UNESCAPTE {
		unescapteAndFormat(jsonPayload)

	} else if formatOption == ESCAPTE {
		escape(jsonPayload)

	} else {
		fmt.Println("Invalid formatting option provided.")
		os.Exit(0)
	}

}

func escape(jsonPayload string) {
	var stringReplacer *strings.Replacer
	var processedJson string

	stringReplacer = strings.NewReplacer("\"", "\\\"")
	processedJson = stringReplacer.Replace(jsonPayload)
	fmt.Print(processedJson)
	os.Exit(1)
}

func unescapteAndFormat(jsonPayload string) {
	var stringReplacer *strings.Replacer
	var processedJson string

	stringReplacer = strings.NewReplacer("\\\"", "\"")

	processedJson = stringReplacer.Replace(jsonPayload)

	var result bytes.Buffer
	err := json.Indent(&result, []byte(processedJson), "", "\t")

	if err != nil {
		fmt.Println("error formatting json ", err.Error())
	}

	formattedJson := result.String()
	fmt.Print(formattedJson)
}
