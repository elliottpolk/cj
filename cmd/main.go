package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/elliottpolk/cj"

	"github.com/pkg/errors"
	cli "gopkg.in/urfave/cli.v2"
)

var (
	version  string
	compiled string = fmt.Sprint(time.Now().Unix())

	inputFileFlag = &cli.StringFlag{
		Name:    "input-file",
		Aliases: []string{"in"},
		Usage:   "file to convert from",
	}

	outputFileFlag = &cli.StringFlag{
		Name:    "output-file",
		Aliases: []string{"out"},
		Usage:   "file to write converted data to",
	}

	delimiterFlag = &cli.StringFlag{
		Name:    "delimiter",
		Aliases: []string{"d"},
		Value:   ",",
		Usage:   "",
	}
)

type row map[string]interface{}

func main() {
	ct, err := strconv.ParseInt(compiled, 10, 64)
	if err != nil {
		panic(err)
	}

	app := cli.App{
		Name:        "cj",
		Copyright:   fmt.Sprintf("Copyright © 2018 - %s Elliott Polk", time.Now().Format("2006")),
		Description: "CSV to JSON converter",
		Version:     version,
		Compiled:    time.Unix(ct, -1),
		Flags: []cli.Flag{
			inputFileFlag,
			outputFileFlag,
			delimiterFlag,
		},
		Action: func(context *cli.Context) error {
			var reader *os.File

			if infile := context.String(inputFileFlag.Name); len(infile) < 1 {
				//  check to see if the CSV was piped in using something like `cat`
				fi, err := os.Stdin.Stat()
				if err != nil {
					return cli.Exit(errors.Wrap(err, "unable to stat stdin"), 1)
				}

				//  if nothing was piped in, just exit
				if fi.Mode()&os.ModeCharDevice != 0 {
					return nil
				}

				reader = os.Stdin
			} else {
				var err error
				reader, err = os.Open(infile)
				if err != nil {
					return cli.Exit(errors.Wrap(err, "unable to open input file"), 1)
				}
				defer reader.Close()
			}

			writer := os.Stdout
			if of := context.String(outputFileFlag.Name); len(of) > 0 {
				out, err := os.OpenFile(of, os.O_WRONLY|os.O_CREATE, 0600)
				if err != nil {
					return cli.Exit(errors.Wrap(err, "unable to open output file"), 1)
				}
				defer out.Close()

				writer = out
			}

			delimiter := context.String(delimiterFlag.Name)

			if err := cj.Convert([]rune(delimiter)[0], bufio.NewReader(reader), writer); err != nil {
				return cli.Exit(err, 1)
			}

			return nil
		},
	}

	app.Run(os.Args)
}
