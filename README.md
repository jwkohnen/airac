# airac: a small Go library for calculating aviation AIRAC cycle dates

[![Apache License v2.0](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0.txt)
[![GoDoc](https://godoc.org/github.com/wjkohnen/airac?status.svg)](https://godoc.org/github.com/wjkohnen/airac)
[![Build Status](https://travis-ci.org/wjkohnen/airac.svg?branch=master)](https://travis-ci.org/wjkohnen/airac)
[![Go Report](https://goreportcard.com/badge/github.com/wjkohnen/airac)](https://goreportcard.com/report/github.com/wjkohnen/airac)
[![codebeat badge](https://codebeat.co/badges/84112bfa-9f47-4bb0-b741-c56441e9fdde)](https://codebeat.co/projects/github-com-wjkohnen-airac)
[![codecov](https://codecov.io/gh/wjkohnen/airac/branch/master/graph/badge.svg)](https://codecov.io/gh/wjkohnen/airac)


Package airac provides calculations on Aeronautical Information Regulation And
Control (AIRAC) cycles, i.e. cycle identifiers and effective calendar dates.

Regular, planned Aeronautical Information Publications (AIP) as defined by the
International Civil Aviation Organization (ICAO) are published and become
effective at fixed dates. This package implements the AIRAC cycle definition as
published in the ICAO Aeronautical Information Services Manual (DOC 8126;
AN/872; 6th Edition; 2003). Test cases validate documented dates from 1998 until
2020, including the rare case of a 14th cycle in the year 2020.

A Java port is available at [github.com/wjkohnen/airac-java/](https://github.com/wjkohnen/airac-java/).

Licensed under the Apache License, Version 2.0.
