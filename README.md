# airac: a small Go library for calculating aviation AIRAC cycle dates

[![Apache License v2.0](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0.txt)
[![GoDoc](https://godoc.org/github.com/jwkohnen/airac?status.svg)](https://godoc.org/github.com/jwkohnen/airac)
[![Build Status](https://travis-ci.org/jwkohnen/airac.svg?branch=master)](https://travis-ci.org/jwkohnen/airac)
[![Go Report](https://goreportcard.com/badge/github.com/jwkohnen/airac)](https://goreportcard.com/report/github.com/jwkohnen/airac)
[![codebeat badge](https://codebeat.co/badges/84112bfa-9f47-4bb0-b741-c56441e9fdde)](https://codebeat.co/projects/github-com-jwkohnen-airac)
[![codecov](https://codecov.io/gh/jwkohnen/airac/branch/master/graph/badge.svg)](https://codecov.io/gh/jwkohnen/airac)


Package airac provides calculations on Aeronautical Information Regulation And
Control (AIRAC) cycles, i.e. cycle identifiers and effective calendar dates.

Regular, planned Aeronautical Information Publications (AIP) as defined by the
International Civil Aviation Organization (ICAO) are published and become
effective at fixed dates. This package implements the AIRAC cycle definition as
published in the ICAO Aeronautical Information Services Manual (DOC 8126;
AN/872; 6th Edition; 2003). Test cases validate documented dates from 1998 until
2020, including the rare case of a 14th cycle in the year 2020.


## License

Licensed under the Apache License, Version 2.0.

## See also

A Java port is available at [github.com/jwkohnen/airac-java/](https://github.com/jwkohnen/airac-java/).

## Wikipedia

Article on AIP / AIRAC cycles: https://en.wikipedia.org/wiki/Aeronautical_Information_Publication

There are wiki macros in that article that do basically the same thing as this
library. Though, this library does not trip over the case of 14 cycles per year
(e. g. 1998 and 2020).
