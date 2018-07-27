package main

import (
	"os"

	"gopkg.in/urfave/cli.v2"
)

var (
	version string

	InputFileFlag = &cli.StringFlag{
		Name:    "input-file",
		Aliases: []string{"in"},
		Usage:   "file to convert from",
	}

	OutputFileFlag = &cli.StringFlag{
		Name:    "output-file",
		Aliases: []string{"out"},
		Usage:   "file to write converted data to",
	}

	DelimiterFlag = &cli.StringFlag{
		Name:    "delimiter",
		Aliases: []string{"d"},
		Value:   ",",
		Usage:   "",
	}
)

type row map[string]interface{}

func main() {
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
		Action: csvToJson,
	}

	app.Run(os.Args)
}
