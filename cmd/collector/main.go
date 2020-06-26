package main

import (
	log "github.com/lobsterbandit/wow-splits/internal/logger"
	"github.com/lobsterbandit/wow-splits/pkg/aggregator"
	"github.com/lobsterbandit/wow-splits/pkg/stats"
)

func main() {
	filePaths := aggregator.FindAllSpeedrunSplits("/World of Warcraft")

	characters := make([]*aggregator.Character, 0, len(filePaths))

	for _, path := range filePaths {
		char := aggregator.CreateCharacter(path)
		if char != nil {
			characters = append(characters, char)
		} else {
			log.Logger.Printf("Error parsing account/server/name from %q", path)
		}
	}

	// log found characters

	for _, char := range characters {
		err := char.ParseCharacterData()
		if err != nil {
			log.Logger.Printf("Error parsing character %q: %v", char.SavedVariablesPath, err)
		}
	}

	stats.CalculateStats(characters)
}
