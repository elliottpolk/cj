package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v2"
)

func csvToJson(context *cli.Context) error {
	var in *os.File

	if infile := context.String(InputFileFlag.Names()[0]); len(infile) < 1 {
		//  check to see if the CSV was piped in using something like `cat`
		fi, err := os.Stdin.Stat()
		if err != nil {
			return cli.Exit(errors.Wrap(err, "unable to stat stdin"), 1)
		}

		//  if nothing was piped in, just exit
		if fi.Mode()&os.ModeCharDevice != 0 || fi.Size() < 1 {
			return nil
		}

		in = os.Stdin
	} else {
		var err error
		in, err = os.Open(infile)
		if err != nil {
			return cli.Exit(errors.Wrap(err, "unable to open input file"), 1)
		}
		defer in.Close()
	}

	w := &writer{os.Stdout}
	if outfile := context.String(OutputFileFlag.Names()[0]); len(outfile) > 0 {
		out, err := os.OpenFile(outfile, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return cli.Exit(errors.Wrap(err, "unable to open output file"), 1)
		}
		defer out.Close()

		w.File = out
	}

	reader := csv.NewReader(in)

	delimiter := context.String(DelimiterFlag.Names()[0])
	reader.Comma = []rune(delimiter)[0]

	var columns []string
	for tick := 0; ; tick++ {
		r, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				if err := w.write("]"); err != nil {
					return cli.Exit(errors.Wrap(err, "unable to write closing bracket"), 1)
				}
				return nil
			}

			return cli.Exit(errors.Wrap(err, "unable to read in from input file"), 1)
		}

		if tick != 0 {
			var (
				out []byte
				err error

				m = make(row)
			)

			for i, k := range columns {
				m[k] = r[i]
			}

			out, err = json.Marshal(m)
			if err != nil {
				return cli.Exit(errors.Wrap(err, "unable to convert row to JSON"), 1)
			}

			if tick == 1 {
				if err := w.write("["); err != nil {
					return cli.Exit(errors.Wrap(err, "unable to write opening bracket"), 1)
				}
			}

			if tick > 1 {
				if err := w.write(","); err != nil {
					return cli.Exit(errors.Wrap(err, "unable to write ','"), 1)
				}
			}

			if err := w.write(string(out)); err != nil {
				return cli.Exit(errors.Wrap(err, "unable to write row data"), 1)
			}
		} else {
			columns = r
		}
	}

	return nil
}
