/*
 * Copyright (c) 2019 Johannes Kohnen <jwkohnen-github@ko-sys.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package airac

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
	"time"
)

func TestFromDate(t *testing.T) {
	t.Parallel()
	var airacTests = []struct {
		date    string
		year    int
		ordinal int
	}{
		/*
			ICAO DOC 8126, 6th edition (2003); Paragraph 2.6.2 b):
				the AIRAC effective dates must be in accordance
				with the predetermined, internationally agreed
				schedule of effective dates based on an interval of
				28 days, including 29 January 1998
		*/
		{"1998-01-29", 1998, 2},

		/*
			ICAO DOC 8126, 6th edition (2003); table 2-1 "Schedule of AIRAC effective date 2003-2012"
		*/
		// first airac of the year and its predecessor
		{"2003-01-23", 2003, 1},

		{"2004-01-21", 2003, 13},
		{"2004-01-22", 2004, 1},

		{"2005-01-19", 2004, 13},
		{"2005-01-20", 2005, 1},

		{"2006-01-18", 2005, 13},
		{"2006-01-19", 2006, 1},

		{"2007-01-17", 2006, 13},
		{"2007-01-18", 2007, 1},

		{"2008-01-16", 2007, 13},
		{"2008-01-17", 2008, 1},

		{"2009-01-14", 2008, 13},
		{"2009-01-15", 2009, 1},

		{"2010-01-13", 2009, 13},
		{"2010-01-14", 2010, 1},

		{"2011-01-12", 2010, 13},
		{"2011-01-13", 2011, 1},

		{"2012-01-11", 2011, 13},
		{"2012-01-12", 2012, 1},

		/*
			http://www.eurocontrol.int/articles/airac-adherence-monitoring-phase-1-p-03
		*/
		// first airac of the year and its predecessor (2013 2020)
		{"2013-01-09", 2012, 13},
		{"2013-01-10", 2013, 1},

		{"2014-01-08", 2013, 13},
		{"2014-01-09", 2014, 1},

		{"2015-01-07", 2014, 13},
		{"2015-01-08", 2015, 1},

		{"2016-01-06", 2015, 13},
		{"2016-01-07", 2016, 1},

		{"2017-01-04", 2016, 13},
		{"2017-01-05", 2017, 1},

		{"2018-01-03", 2017, 13},
		{"2018-01-04", 2018, 1},

		{"2019-01-02", 2018, 13},
		{"2019-01-03", 2019, 1},

		{"2020-01-01", 2019, 13},
		{"2020-01-02", 2020, 1},

		// the 'special' one:
		{"2020-12-30", 2020, 13},
		{"2020-12-31", 2020, 14},

		// calculated manually
		{"2021-01-27", 2020, 14},
		{"2021-01-28", 2021, 1},
		{"2003-01-22", 2002, 13},
		{"1964-01-16", 1964, 1},
		{"1901-01-10", 1901, 1},
		{"1998-12-31", 1998, 14},

		{"1963-12-31", 1963, 13},
	}

	for _, testc := range airacTests {
		testc := testc
		t.Run(testc.date, func(t *testing.T) {
			t.Parallel()
			tdate, _ := time.Parse(format, testc.date)
			got := FromDate(tdate)
			if got.Year() != testc.year || got.Ordinal() != testc.ordinal {
				t.Errorf("Date %s (%s): want: %02d%02d, got: %s",
					testc.date, tdate.Weekday(), testc.year%100, testc.ordinal, got.String())
			}
		})
	}
}

func TestNextPrevious(t *testing.T) {
	t.Parallel()
	testt := []struct {
		date        string
		prevyear    int
		prevordinal int
		nextyear    int
		nextordinal int
	}{
		{"2006-01-20", 2005, 13, 2006, 2},
		{"2021-01-01", 2020, 13, 2021, 1},
	}
	for _, testc := range testt {
		tdate, _ := time.Parse(format, testc.date)
		got := FromDate(tdate)
		gotPrev := got - 1
		gotNext := got + 1

		if gotPrev.Year() != testc.prevyear {
			t.Errorf("got %v, want %v", gotPrev.Year(), testc.prevyear)
		}
		if gotPrev.Ordinal() != testc.prevordinal {
			t.Errorf("got %v, want %v", gotPrev.Ordinal(), testc.prevordinal)
		}
		if gotNext.Year() != testc.nextyear {
			t.Errorf("got %v, want %v", gotNext.Year(), testc.nextyear)
		}
		if gotNext.Ordinal() != testc.nextordinal {
			t.Errorf("got %v, want %v", gotNext.Ordinal(), testc.nextordinal)
		}
	}
}

func TestFromString(t *testing.T) {
	t.Parallel()
	testt := []struct {
		airac     string
		effective string
		year      int
		ordinal   int
		valid     bool
	}{
		{"2014", "2020-12-31", 2020, 14, true},
		{"1511", "2015-10-15", 2015, 11, true},
		{"1514", "", 0, 0, false},
		{"1501", "2015-01-08", 2015, 1, true},
		{"9999", "", 0, 0, false},
		{"6401", "1964-01-16", 1964, 1, true},
		{"6301", "2063-01-04", 2063, 1, true},
		{"6313", "2063-12-06", 2063, 13, true},
		{"9913", "1999-12-30", 1999, 13, true},
		{"09913", "", 0, 0, false},
		{"101", "", 0, 0, false},
		{"0000", "", 0, 0, false},
		{"160a", "", 0, 0, false},
		{"1a01", "", 0, 0, false},
		{"1016", "", 0, 0, false},
		{"10-1", "", 0, 0, false},
		{"-101", "", 0, 0, false},
		{"", "", 0, 0, false},
		{"nope", "", 0, 0, false},
		{"0", "", 0, 0, false},
		{"-0", "", 0, 0, false},
		{"-1", "", 0, 0, false},
		{"11", "", 0, 0, false},
		{"011", "", 0, 0, false},
		{"a", "", 0, 0, false},
		{"aa", "", 0, 0, false},
		{"", "", 0, 0, false},
	}

	for i, testc := range testt {
		i, testc := i, testc
		t.Run(fmt.Sprintf("%02d_\"%s\"", i, testc.airac), func(t *testing.T) {
			t.Parallel()
			got, err := FromString(testc.airac)
			if testc.valid && err != nil {
				t.Errorf("AIRAC \"%v\" did not parse: %v", testc.airac, err)
				return
			}
			if !testc.valid && err == nil {
				t.Errorf("AIRAC \"%v\" parsed to %v, but should have raised an error!!", testc.airac, got)
				return
			}
			if !testc.valid && err != nil {
				t.Logf("Test string \"%s\" rightfully yields error: %v", testc.airac, err)
				return
			}

			wantEffective, err := time.Parse(format, testc.effective)
			if err != nil {
				t.Fatalf("test case broken: %v", err)
			}
			if got.Year() != testc.year ||
				got.Ordinal() != testc.ordinal ||
				!got.Effective().Equal(wantEffective) {
				t.Errorf("AIRAC \"%v\", want %02d%02d (eff: %v), got %v",
					testc.airac, testc.year%100, testc.ordinal, testc.effective, got.LongString())
			}
		})
	}
}

func TestFromStringZeroOrdinal(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		i := i
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			s := fmt.Sprintf("%02d00", i)
			got, err := FromString(s)
			if err == nil {
				t.Errorf("Identifier %s yields AIRAC cycle %s.", s, got)
				return
			}
			t.Logf("Identifier %s rightfully returns error: %v", s, err)
		})
	}
}

func TestFromStringMust(t *testing.T) {
	t.Parallel()
	got := FromStringMust("1201")
	want := "1201"
	if got.String() != want {
		t.Errorf("Got %s, want %s", got, want)
	}

	func() {
		invalid := "1614"
		defer func() {
			r := recover()
			if r == nil {
				t.Errorf("TestFromStringMust(\"%s\") should have paniced, but didn't.", invalid)
			} else {
				t.Logf("TestFromStringMust(\"%s\") rightfully paniced: %v", invalid, r)
			}
		}()

		FromStringMust(invalid)
	}()
}

func TestOverflow(t *testing.T) {
	t.Parallel()
	// there will be an overflow after April 4th, 2193
	last := time.Date(2193, time.April, 4, 0, 0, 0, 0, time.UTC)
	for a := AIRAC(1); a.Effective().Before(last); a++ {
		prev := a - 1
		diff := a.Effective().Sub(prev.Effective())
		if diff != 28*24*time.Hour {
			t.Errorf("Time difference between cycle (%s) %d/%02d and (%s) %d/%02d wrong, want 28 days, got %s.",
				a.LongString(), a.Year(), a.Ordinal(), prev.LongString(), prev.Year(), prev.Ordinal(), diff)
			break
		}
	}
}

func TestLastAiracOfYear(t *testing.T) {
	t.Parallel()
	for year := epoch.Year(); year < 2193; year++ {
		year := year
		t.Run(strconv.Itoa(year), func(t *testing.T) {
			t.Parallel()
			airac := FromDate(time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC))
			if airac.Year() != year {
				t.Errorf("want %d, got %d", year, airac.Year())
			}
		})
	}
}

func ExampleFromDate() {
	shalom := time.Date(2012, time.August, 26, 0, 0, 0, 0, time.UTC)
	airac := FromDate(shalom)

	fmt.Printf("At %s the current AIRAC cycle was %s.\n\n",
		shalom.Format("2006-01-02"),
		airac.LongString(),
	)

	fmt.Printf("Short identifier: %s", airac)

	// Output:
	// At 2012-08-26 the current AIRAC cycle was 1209 (effective: 2012-08-23; expires: 2012-09-19).
	//
	// Short identifier: 1209

}

func BenchmarkFromString(b *testing.B) {
	r := make([]AIRAC, 0, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = append(r, FromStringMust("2014"))
	}
	_ = r
}

func BenchmarkFromDate(b *testing.B) {
	r := make([]AIRAC, 0, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = append(r, FromDate(time.Now()))
	}
	_ = r
}

func TestTypeAlias(t *testing.T) {
	t.Parallel()
	lower := Airac(42)
	upper := AIRAC(42)
	if lower != upper {
		t.Error("Old type does not equal new type!")
	}
}

func ExampleByChrono() {
	airacs := []AIRAC{
		FromStringMust("1213"),
		FromStringMust("1201"),
		FromStringMust("1207"),
	}
	fmt.Println("Not Sorted:    ", airacs)

	sort.Sort(ByChrono(airacs))
	fmt.Println("Sorted:        ", airacs)

	sort.Sort(sort.Reverse(ByChrono(airacs)))
	fmt.Println("Sorted reverse:", airacs)

	// Output:
	// Not Sorted:     [1213 1201 1207]
	// Sorted:         [1201 1207 1213]
	// Sorted reverse: [1213 1207 1201]
}
