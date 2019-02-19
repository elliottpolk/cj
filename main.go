package main

import (
	"log"
	"os"

	"gopkg.in/urfave/cli.v2"
)

var version string

const inputFileFlagName = "input-file"
const outputFileFlagName = "output-file"
const delimiterFlagName = "delimiter"

type row map[string]interface{}

func main() {
	InputFileFlag := &cli.StringFlag{
		Name:    inputFileFlagName,
		Aliases: []string{"in"},
		Usage:   "file to convert from",
	}

	OutputFileFlag := &cli.StringFlag{
		Name:    outputFileFlagName,
		Aliases: []string{"out"},
		Usage:   "file to write converted data to",
	}

	DelimiterFlag := &cli.StringFlag{
		Name:    delimiterFlagName,
		Aliases: []string{"d"},
		Value:   ",",
		Usage:   "",
	}

	app := cli.App{
		Name:        "cj",
		Copyright:   "Copyright © 2018 Elliott Polk",
		Description: "CSV to JSON converter",
		Version:     version,
		Flags: []cli.Flag{
			InputFileFlag,
			OutputFileFlag,
			DelimiterFlag,
		},
		Action: csvToJSON,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("app.Run: %v", err)
	}
}
