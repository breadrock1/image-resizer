package cmd

import (
	"encoding/json"
	"log"
	"os"
)

var (
	release   = "0.0.1"
	buildDate = "2024-05-23"
	gitHash   = ""
)

type versionInfo struct {
	Release   string
	BuildDate string
	GitHash   string
}

func PrintVersion() {
	encoder := json.NewEncoder(os.Stdout)
	versionObject := versionInfo{
		Release:   release,
		BuildDate: buildDate,
		GitHash:   gitHash,
	}

	if err := encoder.Encode(versionObject); err != nil {
		log.Fatalf("error while decode version info: %v\n", err)
	}
}
