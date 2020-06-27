package main

import (
	log "github.com/lobsterbandit/wow-splits/internal/logger"
	"github.com/lobsterbandit/wow-splits/pkg/character"
	"github.com/lobsterbandit/wow-splits/pkg/file"
	"github.com/lobsterbandit/wow-splits/pkg/stats"
)

func main() {
	filePaths := file.FindAllFiles("/World of Warcraft")

	characters := make([]*character.Character, 0, len(filePaths))

	for _, path := range filePaths {
		char := character.CreateCharacter(path)
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
