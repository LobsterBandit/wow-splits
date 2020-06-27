package main

import (
	"flag"
	"os"

	log "github.com/lobsterbandit/wow-splits/internal/logger"
	"github.com/lobsterbandit/wow-splits/pkg/character"
	"github.com/lobsterbandit/wow-splits/pkg/file"
)

func main() {
	wowDir := flag.String("wowdir", "", "(Required) path to \"World of Warcraft\" install `directory`")
	printTree := flag.Bool("tree", false, "print Account/Server/Character tree")
	printTreeWithTimes := flag.Bool("tree-with-times", false, "print Account/Server/Character tree with per level times")
	debug := flag.Bool("debug", false, "print debug logs to stdout")

	flag.Parse()

	log.CreateGlobalLogger(*debug)

	if *wowDir == "" {
		flag.Usage()
		os.Exit(1)
	}

	log.Logger.Println("Collecting character leveling times...")

	filePaths := file.FindAllFiles(*wowDir, *debug)

	characters := make([]*character.Character, 0, len(filePaths))

	for _, path := range filePaths {
		char := character.CreateCharacter(path)
		if char != nil {
			characters = append(characters, char)
		} else {
			log.Logger.Printf("Error parsing account/server/name from %q", path)
		}
	}

	for _, char := range characters {
		err := char.ParseCharacterData(*debug)
		if err != nil {
			log.Logger.Printf("Error parsing character %q: %v", char.SavedVariablesPath, err)
		}
	}

	if *printTree || *printTreeWithTimes {
		character.PrintCharacterTree(characters, *printTreeWithTimes)
	}
}
