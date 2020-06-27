package character

import (
	"regexp"
)

type Level struct {
	Level         int
	LevelTime     int
	AggregateTime int
}

type Character struct {
	Account            string
	Server             string
	Name               string
	Class              string
	SavedVariablesPath string
	Times              map[int]*Level
}

var charMetadataRegexp = regexp.MustCompile(`_classic_/WTF/Account/(?P<AccountName>\w+)/(?P<ServerName>\w+)/(?P<CharacterName>\w+)`)

func CreateCharacter(path string) *Character {
	matches := charMetadataRegexp.FindStringSubmatch(path)

	if matches == nil {
		return nil
	}

	return &Character{
		Account:            matches[1],
		Server:             matches[2],
		Name:               matches[3],
		SavedVariablesPath: path,
		Times:              make(map[int]*Level),
	}
}
