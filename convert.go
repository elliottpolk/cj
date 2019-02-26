package cj

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// Convert is a simple wrapper that takes in an io.Reader and io.Writer, converting the CSV data from the
// reader to simple JSON and outputting to the supplied writer
func Convert(delimiter rune, in io.Reader, out io.Writer) error {
	reader := csv.NewReader(in)
	reader.Comma = delimiter
	reader.TrimLeadingSpace = true

	var columns []string
	for tick := 0; ; tick++ {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				if _, err := fmt.Fprint(out, "]"); err != nil {
					return errors.Wrap(err, "unable to write closing bracket")
				}
				return nil
			}

			return errors.Wrap(err, "unable to read from input")
		}

		// assumes header starts at line 0
		if tick == 0 {
			columns = line
			continue
		}

		row := make(map[string]interface{})

		for i, key := range columns {
			row[key] = line[i]
		}

		data, err := json.Marshal(row)
		if err != nil {
			return errors.Wrap(err, "unable to convert row to JSON")
		}

		// first row so start with '[' since it should always be a list
		if tick == 1 {
			if _, err := fmt.Fprint(out, "["); err != nil {
				return errors.Wrap(err, "unable to write opening bracket")
			}
		}

		if tick > 1 {
			if _, err := fmt.Fprint(out, ","); err != nil {
				return errors.Wrap(err, "unable to write ','")
			}
		}

		if _, err := fmt.Fprint(out, string(data)); err != nil {
			return errors.Wrap(err, "unable to write row data")
		}
	}
}
