package aggregator

import (
	"os"
	"path/filepath"
	"strings"

	log "github.com/lobsterbandit/wow-splits/pkg/logger"
)

const SpeedrunSplitsFile string = "SpeedrunSplits.lua"

func FindAllSpeedrunSplits(wowDir string) (files []string) {
	log.Logger.Println("WoW dir = ", wowDir)

	err := filepath.Walk(wowDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// skip global savedvariables fille
		if info.Name() == SpeedrunSplitsFile &&
			len(strings.Split(path, string(filepath.Separator))) == 10 {
			log.Logger.Printf("Found %q at %q", info.Name(), path)
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		log.Logger.Printf("Error walking the path %q: %v", wowDir, err)

		return files
	}

	return files
}
