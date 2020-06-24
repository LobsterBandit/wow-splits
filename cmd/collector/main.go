package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/lobsterbandit/wow-splits/pkg/aggregator"
	"github.com/lobsterbandit/wow-splits/pkg/stats"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

func main() {
	filePaths := aggregator.FindAllSpeedrunSplits("/World of Warcraft")

	characters := make([]*aggregator.Character, 0, 10)

	for i, file := range filePaths {
		data, err := aggregator.ParseCharacter(file)
		if err != nil {
			logger.Printf("Error reading %q: %v", file, err)
		}

		pretty, _ := json.MarshalIndent(data, "", "\t")

		logger.Printf("%v: %s", i, pretty)

		characters = append(characters, data)
	}

	stats.CalculateStats(characters)
}
