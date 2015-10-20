# airac: a small Go library for calculating aviation AIRAC cycle dates

[![GNU Affero Public License v3](https://www.gnu.org/graphics/agplv3-88x31.png)](https://www.gnu.org/licenses/agpl-3.0.html)
[![GoDoc](https://godoc.org/github.com/wjkohnen/airac?status.svg)](https://godoc.org/github.com/wjkohnen/airac)
[![Build Status](https://travis-ci.org/wjkohnen/airac.svg?branch=master)](https://travis-ci.org/wjkohnen/airac)

Package airac provides calculations on Aeronautical Information Regulation And
Control (AIRAC) cycles, i.e. cycle identifiers and effective calendar dates.

Regular, planned Aeronautical Information Publications (AIP) as defined by the
International Civil Aviation Organization (ICAO) are published and become
effective at fixed dates. This package implements the AIRAC cycle definition as
published in the ICAO Aeronautical Information Services Manual (DOC 8126;
AN/872; 6th Edition; 2003). Test cases validate documented dates from 1998 until
2020, including the rare case of a 14th cycle in the year 2020.

Licensed under GNU Affero General Public License version 3.0.
