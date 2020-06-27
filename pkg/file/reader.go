package file

import (
	"io/ioutil"

	log "github.com/lobsterbandit/wow-splits/internal/logger"
)

// Table Keys
// Global:
// - SpeedrunSplitsGold
// - SpeedrunSplitsPB
// - SpeedrunSplitsOptions
// Character:
// - SpeedrunSplits

func ReadFile(path string, debug bool) (data string, err error) {
	if debug {
		log.Logger.Printf("Reading %q", path)
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Logger.Printf("Error reading %q: %v", path, err)
		return "", err
	}

	return string(content), nil
}
