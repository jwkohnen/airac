/*
 * Copyright (c) 2020 Johannes Kohnen <jwkohnen-github@ko-sys.com>
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

/*
Package airac provides calculations on Aeronautical Information Regulation And
Control (AIRAC) cycles, i.e. cycle identifiers and effective calendar dates.

Regular, planned Aeronautical Information Publications (AIP) as defined by the
International Civil Aviation Organization (ICAO) are published and become
effective at fixed dates. This package implements the AIRAC cycle definition as
published in the ICAO Aeronautical Information Services Manual (DOC 8126;
AN/872; 6th Edition; 2003). Test cases validate documented dates from 1998 until
2020, including the rare case of a 14th cycle in the year 2020.

Licensed under the Apache License, Version 2.0.
*/
package airac

// nolint:godox
func _() {
	/* BUG(jwkohnen): The two digit year identifier of the FromString method will
	   interpret the year as between 1964 and 2063. Other methods than FromString do
	   not show this range restriction. This time window is more or less arbitrary and
	   may change. */

	/* BUG(jwkohnen): This package assumes that AIRAC cycles are effective from
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

	/* BUG(jwkohnen): Calculations that include calendar dates before the internal
	   epoch (1901-01-10; 63 years before the AIRAC system was introduced by the ICAO)
	   and after year 2192 may silently produce wrong data. */

	/* BUG(jwkohnen): This package only provides calculations on effective dates,
	   not publication or reception dates etc. Although effective dates are clearly
	   defined and are consistent at least between 1998 until 2020, the derivative
	   dates changed historically.[citation needed] */
}
