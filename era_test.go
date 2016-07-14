package era_jp

import (
	"testing"
)

var testcases = []struct {
	year int
	era  string
}{
	{1900, `明治`},
	{1911, `明治`},
	{1912, `大正`},
	{1925, `大正`},
	{1926, `昭和`},
	{1988, `昭和`},
	{1989, `平成`},
	{2016, `平成`},
}

func TestSimple(t *testing.T) {
	for _, testcase := range testcases {
		year := testcase.year
		got := ToEra(year)
		expected := testcase.era
		if got != expected {
			t.Fatalf("Expected %v for year %v, but %v:", expected, year, got)
		}
	}
}
