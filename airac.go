/*
 * Copyright (c) 2016 Wolfgang Johannes Kohnen <wjkohnen@users.noreply.github.com>
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
	"regexp"
	"strconv"
	"time"
)

// Airac represents an AIRAC cycle.
type Airac int

// Year returns the year for this AIRAC cycle's identifier.
func (a Airac) Year() int {
	return a.Effective().Year()
}

// Ordinal returns the ordinal for this AIRAC cycle's identifier.
func (a Airac) Ordinal() int {
	return (a.Effective().YearDay()-1)/28 + 1
}

// Effective returns the effective date of this AIRAC cycle.
func (a Airac) Effective() time.Time {
	return epoch.Add(time.Duration(a) * durationCycle)
}

// LongString returns a verbose representation of this AIRAC cycle.
// "YYOO (effective: YYYY-MM-DD; expires: YYYY-MM-DD)"
func (a Airac) LongString() string {
	n := a + 1
	return fmt.Sprintf("%02d%02d (effective: %s; expires: %s)",
		a.Year()%100,
		a.Ordinal(),
		a.Effective().Format(format),
		n.Effective().Add(-1).Format(format),
	)
}

// String returns a short representation of this AIRAC cycle. "YYOO"
func (a Airac) String() string {
	return fmt.Sprintf("%02d%02d", a.Year()%100, a.Ordinal())
}

// FromDate returns the AIRAC cycle that occurred at date. A date before the
// internal epoch (1901-01-10) may return wrong data. The upper limit is
// year 2192.
func FromDate(date time.Time) Airac {
	a := date.Sub(epoch) / durationCycle
	return Airac(a)
}

// FromString returns an AIRAC cycle that matches the identifier <yyoo>,
// i.e. the last two digits of the year and the ordinal, each with leading
// zeros. This works for years between 1964 and 2063. Identifiers between
// "6401" and "9913" are interpreted as AIRAC cycles between the years 1964
// and 1999 inclusive. AIRAC cycles between "0001" and "6313" are
// interpreted as AIRAC cycles between the years 2000 and 2063 inclusive.
func FromString(yyoo string) (Airac, error) {
	year, ordinal, err := parseIdentifier(yyoo)
	if err != nil {
		return -1, err
	}

	lastAiracOfPreviousYear := FromDate(time.Date(year-1, time.December, 31, 0, 0, 0, 0, time.UTC))
	airac := lastAiracOfPreviousYear + Airac(ordinal)

	if airac.Year() != year {
		return -1, fmt.Errorf("year %d does not have %d airac cycles", year, ordinal)
	}

	return airac, nil
}

func parseIdentifier(yyoo string) (year, ordinal int, err error) {
	m := identifierRegex.FindStringSubmatch(yyoo)
	if m == nil {
		return -1, -1, fmt.Errorf("illegal airac identifier: %s", yyoo)
	}
	year, _ = strconv.Atoi(m[1])
	ordinal, _ = strconv.Atoi(m[2])

	if year > 63 {
		year = 1900 + year
	} else {
		year = 2000 + year
	}

	return year, ordinal, nil
}

const (
	format = "2006-01-02"
)

var (
	epoch           = time.Date(1901, time.January, 10, 0, 0, 0, 0, time.UTC)
	durationCycle   = 28 * 24 * time.Hour
	identifierRegex = regexp.MustCompile(`^(\d{2})(\d{2})$`)
)
