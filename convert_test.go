package cj

import (
	"bytes"
	"fmt"
	"io"
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

		if got := buf.Bytes(); !bytes.Equal(want, got) {
			t.Errorf("\nwant\t%s\ngot\t%s", string(want), string(got))
		}

	}
}

func BenchmarkConvert(b *testing.B) {
	csr := testCSVSampleReader()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		if err := Convert(',', csr, &buf); err != nil {
			b.Fatal(err)
		}
	}
}

func testCSVSampleReader() io.Reader {
	return bytes.NewBufferString(
		`"LatD", "LatM", "LatS", "NS", "LonD", "LonM", "LonS", "EW", "City", "State"
	41,    5,   59, "N",     80,   39,    0, "W", "Youngstown", OH
	42,   52,   48, "N",     97,   23,   23, "W", "Yankton", SD
	46,   35,   59, "N",    120,   30,   36, "W", "Yakima", WA
	42,   16,   12, "N",     71,   48,    0, "W", "Worcester", MA
	43,   37,   48, "N",     89,   46,   11, "W", "Wisconsin Dells", WI
	36,    5,   59, "N",     80,   15,    0, "W", "Winston-Salem", NC
	49,   52,   48, "N",     97,    9,    0, "W", "Winnipeg", MB
	39,   11,   23, "N",     78,    9,   36, "W", "Winchester", VA
	34,   14,   24, "N",     77,   55,   11, "W", "Wilmington", NC
	39,   45,    0, "N",     75,   33,    0, "W", "Wilmington", DE
	48,    9,    0, "N",    103,   37,   12, "W", "Williston", ND`)
}
