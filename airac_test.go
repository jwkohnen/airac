/*
 * Copyright (C) 2015 Wolfgang Johannes Kohnen <wjkohnen@users.noreply.github.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package airac

import (
	"fmt"
	"testing"
	"time"
)

func TestFromDate(t *testing.T) {
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
	}

	for _, test := range airacTests {
		tdate, _ := time.Parse(format, test.date)
		got := FromDate(tdate)
		if got.Year() != test.year || got.Ordinal() != test.ordinal {
			t.Errorf("Date %s (%s): want: %02d%02d, got: %s",
				test.date, tdate.Weekday(), test.year%100, test.ordinal, got.String())
		}
	}
}

func TestNextPrevious(t *testing.T) {
	tests := []struct {
		date        string
		prevyear    int
		prevordinal int
		nextyear    int
		nextordinal int
	}{
		{"2006-01-20", 2005, 13, 2006, 2},
		{"2021-01-01", 2020, 13, 2021, 1},
	}
	for _, test := range tests {
		tdate, _ := time.Parse(format, test.date)
		got := FromDate(tdate)
		gotPrev := got - 1
		gotNext := got + 1

		if gotPrev.Year() != test.prevyear {
			t.Errorf("got %v, want %v", gotPrev.Year(), test.prevyear)
		}
		if gotPrev.Ordinal() != test.prevordinal {
			t.Errorf("got %v, want %v", gotPrev.Ordinal(), test.prevordinal)
		}
		if gotNext.Year() != test.nextyear {
			t.Errorf("got %v, want %v", gotNext.Year(), test.nextyear)
		}
		if gotNext.Ordinal() != test.nextordinal {
			t.Errorf("got %v, want %v", gotNext.Ordinal(), test.nextordinal)
		}
	}
}

func TestFromString(t *testing.T) {
	tests := []struct {
		airac     string
		effective string
		year      int
		ordinal   int
		ok        bool
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
		{"101", "", 0, 0, false},
		{"160a", "", 0, 0, false},
		{"1a01", "", 0, 0, false},
	}

	for _, test := range tests {
		got, err := FromString(test.airac)
		if test.ok && err != nil {
			t.Errorf("AIRAC \"%v\" did not parse: %v", test.airac, err)
			continue
		}
		if !test.ok && err == nil {
			t.Errorf("AIRAC \"%v\" parsed to %v, but should have raised an error!!", test.airac, got)
			continue
		}
		if !test.ok && err != nil {
			continue
		}
		wantEffective, err := time.Parse(format, test.effective)
		if err != nil {
			t.Fatalf("test case broken: %v", err)
		}
		if got.Year() != test.year ||
			got.Ordinal() != test.ordinal ||
			!got.Effective().Equal(wantEffective) {
			t.Errorf("AIRAC \"%v\", want %02d%02d (eff: %v), got %v",
				test.airac, test.year%100, test.ordinal, test.effective, got.LongString())
		}
	}
}

func TestOverflow(t *testing.T) {
	// there will be an overflow after April 4th, 2193
	last := time.Date(2193, time.April, 4, 0, 0, 0, 0, time.UTC)
	for a := Airac(1); a.Effective().Before(last); a++ {
		prev := a - 1
		diff := a.Effective().Sub(prev.Effective())
		if diff != 28*24*time.Hour {
			t.Errorf("Time difference between cycle (%s) %d/%02d and (%s) %d/%02d wrong, want 28 days, got %s.",
				a.LongString(), a.Year(), a.Ordinal(), prev.LongString(), prev.Year(), prev.Ordinal(), diff)
			break
		}
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
