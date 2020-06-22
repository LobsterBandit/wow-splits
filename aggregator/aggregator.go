package aggregator

import (
	"fmt"
	"os"
	"path/filepath"
)

const SpeedrunSplitsFile string = "SpeedrunSplits.lua"

func HelloAggregator() {
	fmt.Println("Hello from aggregator!")
}

func FindAllSpeedrunSplits(wowDir string) (files []string) {
	fmt.Println("WoW dir =", wowDir)

	err := filepath.Walk(wowDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == SpeedrunSplitsFile {
			fmt.Printf("found %q at %q\n", info.Name(), path)
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", wowDir, err)

		return files
	}

	return files
}
