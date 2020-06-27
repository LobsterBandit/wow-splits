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

func ReadFile(path string) (data string, err error) {
	log.Logger.Printf("Reading %q", path)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Logger.Printf("Error reading %q: %v", path, err)
		return "", err
	}

	return string(content), nil
}
