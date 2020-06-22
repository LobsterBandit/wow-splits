package aggregator

import (
	"fmt"
	"io/ioutil"
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
			fmt.Printf("\nfound %q at %q", info.Name(), path)
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("\nerror walking the path %q: %v", wowDir, err)

		return files
	}

	return files
}

func ReadSpeedrunSplits(path string) (data string, err error) {
	fmt.Printf("\nReading file at %q", path)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("\nerror reading %q: %v", path, err)
		return "", err
	}

	return string(content), nil
}
