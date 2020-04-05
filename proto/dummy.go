package proto

/*
dummy file that fixes this error when build tag proto is not specified:
	$ go mod tidy
	[...]
	github.com/jwkohnen/airac imports
	github.com/jwkohnen/airac/proto: module github.com/jwkohnen/airac@latest found (v1.0.4), but does not contain package github.com/jwkohnen/airac/proto
*/
