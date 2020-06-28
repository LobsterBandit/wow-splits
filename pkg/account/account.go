package account

import (
	"path/filepath"

	"github.com/lobsterbandit/wow-splits/pkg/file"
)

type Account struct {
	Name               string
	Path               string
	SavedVariablesPath string
	Servers            map[string][]*Character
	ServerCount        int
	Characters         []*Character
	CharacterCount     int
}

func CreateAllAccounts(wowDir string, debug bool) (accounts []*Account, err error) {
	accountDirs, err := file.FindAllAccountPaths(wowDir, debug)
	if err != nil {
		return nil, err
	}

	for _, dir := range accountDirs {
		account := &Account{
			Name:               filepath.Base(dir),
			Path:               dir,
			SavedVariablesPath: filepath.Join(dir, "SavedVariables", file.SpeedrunSplitsFile),
			Servers:            make(map[string][]*Character),
		}

		accounts = append(accounts, account)
	}

	return
}

func (a *Account) PopulateAccountData(debug bool) error {
	dataPaths, err := file.FindAllDataPaths(a.Path, debug)
	if err != nil {
		return err
	}

	if existingCharacters := a.Characters; existingCharacters == nil {
		a.Characters = make([]*Character, 0, 20)
	}

	for _, path := range dataPaths {
		if path == a.SavedVariablesPath {
			continue
		}

		character := CreateCharacter(path)

		err := character.ParseCharacterData(debug)
		if err != nil {
			return err
		}

		a.Characters = append(a.Characters, character)
		a.CharacterCount++

		if existingServer := a.Servers[character.Server]; existingServer == nil {
			a.Servers[character.Server] = make([]*Character, 0, 20)
			a.ServerCount++
		}
		a.Servers[character.Server] = append(a.Servers[character.Server], character)
	}

	return nil
}
