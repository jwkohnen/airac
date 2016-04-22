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

/*
Package airac provides calculations on Aeronautical Information Regulation And
Control (AIRAC) cycles, i.e. cycle identifiers and effective calendar dates.

Regular, planned Aeronautical Information Publications (AIP) as defined by the
International Civil Aviation Organization (ICAO) are published and become
effective at fixed dates. This package implements the AIRAC cycle definition as
published in the ICAO Aeronautical Information Services Manual (DOC 8126;
AN/872; 6th Edition; 2003). Test cases validate documented dates from 1998 until
2020, including the rare case of a 14th cycle in the year 2020.

Licensed under GNU Affero General Public License version 3.0.
*/
package airac

/* BUG(wjkohnen): The two digit year identifier of the FromString method will
interpret the year as between 1964 and 2063. Other methods than FromString do
not show this range restriction. This time window is more or less arbitrary and
may change. */

/* BUG(wjkohnen): This package assumes that AIRAC cycles are effective from
the effective date at 00:00:00 UTC until 27 days later at 23:59:59.999999999
UTC. That is not correct:

ICAO DOC 8126, 6th Edition (2003), paragraph 2.6.4:
  "In addition to the use of a predetermined schedule of effective AIRAC dates,
  Coordinated Universal Time (UTC) must also be used to indicate the time when
  the AIRAC information will become effective. Since Annex 15, paragraph 3.2.3
  specifies that the Gregorian calendar and UTC must be used as the temporal
  reference system for international civil aviation, in addition to AIRAC
  dates, 00:01 UTC must be used to indicate the time when the AIRAC-based
  information will become effective."

However I won't "fix" this, because that may just confuse users. */

/* BUG(wjkohnen): Calculations that include calendar dates before the internal
epoch (1901-01-10; 63 years before the AIRAC system was introduced by the ICAO)
and after year 2192 may silently produce wrong data. */

/* BUG(wjkohnen): This package only provides calculations on effective dates,
not publication or reception dates etc. Although effective dates are clearly
defined and are consistent at least between 1998 until 2020, the derivative
dates changed historically.[citation needed] */
