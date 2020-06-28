package account

import (
	"path/filepath"

	"github.com/lobsterbandit/wow-splits/pkg/file"
)

type WowInstall struct {
	RootDir  string
	Accounts []*Account
}

type Account struct {
	Name               string
	Path               string
	SavedVariablesPath string
	Servers            map[string][]*Character
	ServerCount        int
	Characters         []*Character
	CharacterCount     int
}

func (w *WowInstall) CreateAllAccounts(debug bool) (err error) {
	accountDirs, err := file.FindAllAccountPaths(w.RootDir, debug)
	if err != nil {
		return err
	}

	for _, dir := range accountDirs {
		account := &Account{
			Name:               filepath.Base(dir),
			Path:               dir,
			SavedVariablesPath: filepath.Join(dir, "SavedVariables", file.SpeedrunSplitsFile),
			Servers:            make(map[string][]*Character),
		}

		w.Accounts = append(w.Accounts, account)
	}

	return nil
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
