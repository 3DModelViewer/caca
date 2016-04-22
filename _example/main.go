package main

import (
	"github.com/modelhub/caca"
	"github.com/robsix/golog"
	"fmt"
)

const (
	ccHost = "http://shflinux03:4000"
)

func main() {
	log := golog.NewConsoleLog(0)
	ccClient := caca.NewCacaClient(ccHost, log)

	/*
	structural := "urn:adsk.viewing:fs.file:dXJuOmFkc2sub2JqZWN0czpvcy5vYmplY3Q6bW9kZWxodWJfMDE5MWE5MTc0OTM4NDY0ZjI1OGRmZjUzODE5YjQyNWE4Ny84MjUyNTUyZWRlNmQ0MGJhYTE4ZDcwMmE0Y2FmYjIyMC5ud2Q=/output/0/0.svf"
	hvac := "urn:adsk.viewing:fs.file:dXJuOmFkc2sub2JqZWN0czpvcy5vYmplY3Q6bW9kZWxodWJfMDE5MWE5MTc0OTM4NDY0ZjI1OGRmZjUzODE5YjQyNWE4Ny9jMjc4OGQ4N2U4N2Y0Yzc3YTEzNWRjNTliZGU0ZmQ0YS5ud2Q=/output/0/0.svf"

	data, err := ccClient.RegisterSheet(structural)
	if err != nil {
		log.Critical("%v", err)
	}
	if data != nil {
		str, _ := data.ToString()
		log.Info("%v", str)
	}

	data, err =ccClient.RegisterSheet(hvac)
	if err != nil {
		log.Critical("%v", err)
	}
	if data != nil {
		str, _ := data.ToString()
		log.Info("%v", str)
	}
	*/

	/*
	structuralRegId := "a37c042b-3ac7-438a-bdad-2a610e31017d"
	hvacRegId := "37aa6167-8381-4a61-8a1e-b328c6fc59ce"

	data, err := ccClient.RegisterClashTest(structuralRegId, hvacRegId)

	if err != nil {
		log.Critical("%v", err)
	}
	if data != nil {
		str, _ := data.ToString()
		log.Info("%v", str)
	}
	*/

	clashTestId := "65adf44b-f725-45f1-95e8-98885782a86b"
	data, err := ccClient.GetClashTest(clashTestId)

	if err != nil {
		log.Critical("%v", err)
	}
	if data != nil {
		str, _ := data.ToString()
		log.Info("%v", str)
	}

	fmt.Scanln()
}
