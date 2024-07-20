package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ESCAPTE   string = "e"
	UNESCAPTE string = "u"
)

func main() {

	if len(os.Args) > 23 {
		log.Fatal("Expecting 2 arguments only got ", len(os.Args))
	}

	var jsonPayload string
	flag.StringVar(&jsonPayload, "p", "", "Json payload")

	var formatOption string
	flag.StringVar(&formatOption, "f", "n", "Formatting option")

	flag.Parse()

	if jsonPayload == "" {
		log.Fatal("Json payload is requied but not supplied.")
	}

	var stringReplacer *strings.Replacer

	if formatOption == UNESCAPTE {
		stringReplacer = strings.NewReplacer("\\\"", "\"")
	} else if formatOption == ESCAPTE {
		stringReplacer = strings.NewReplacer("\"", "\\\"")
	} else {
		log.Fatal("Invalid formatting option provided.")
		os.Exit(0)
	}

	var unescapedJson string = stringReplacer.Replace(jsonPayload)
	var result bytes.Buffer
	err := json.Indent(&result, []byte(unescapedJson), "", "\t")

	if err != nil {
		log.Fatal("error formatting json ", err.Error())
	}

	formattedJson := string(result.Bytes())
	fmt.Print(formattedJson)
}
