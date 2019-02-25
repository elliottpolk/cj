package cj

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	for _, src := range []string{"addresses", "airtravel", "cities"} {
		in, err := os.Open(fmt.Sprintf("test_data/%s.csv", src))
		if err != nil {
			t.Fatal(err)
		}
		defer in.Close()

		var buf bytes.Buffer

		if err := Convert(',', in, &buf); err != nil {
			t.Fatal(err)
		}

		want, err := ioutil.ReadFile(fmt.Sprintf("test_data/%s.json", src))
		if err != nil {
			t.Fatal(err)
		}

		if got := buf.Bytes(); bytes.Compare(want, got) != 0 {
			t.Errorf("\nwant\t%s\ngot\t%s", string(want), string(got))
		}

	}
}
