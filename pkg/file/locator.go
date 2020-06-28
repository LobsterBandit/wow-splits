package file

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/lobsterbandit/wow-splits/internal/logger"
)

type CharacterPath struct {
	Server    string
	Character string
}

const SpeedrunSplitsFile string = "SpeedrunSplits.lua"

func FindAllAccountPaths(wowDir string, debug bool) (accountDirs []string, err error) {
	accountPath := filepath.Join(wowDir, "/_classic_/WTF/Account")

	log.Logger.Printf("Searching for accounts in %q...", accountPath)

	files, err := ioutil.ReadDir(accountPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			accountDirs = append(accountDirs, filepath.Join(accountPath, file.Name()))

			if debug {
				log.Logger.Printf("Found account %q", file.Name())
			}
		}
	}

	return
}

func FindAllDataPaths(path string, debug bool) (dataPaths []string, err error) {
	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == SpeedrunSplitsFile {
			dataPaths = append(dataPaths, path)
			if debug {
				log.Logger.Printf("Found data at %q", path)
			}
		}

		return nil
	})

	return
}
