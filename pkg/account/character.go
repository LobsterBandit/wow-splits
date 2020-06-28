package account

import (
	"regexp"
)

type Character struct {
	Server             string
	Name               string
	Class              string
	SavedVariablesPath string
	Level              int
	Times              map[int]*Level
}

type Level struct {
	Level         int
	LevelTime     int
	AggregateTime int
}

var charMetadataRegexp = regexp.MustCompile(`_classic_/WTF/Account/(?P<AccountName>\w+)/(?P<ServerName>\w+)/(?P<CharacterName>\w+)`)

func CreateCharacter(path string) *Character {
	matches := charMetadataRegexp.FindStringSubmatch(path)

	if matches == nil {
		return nil
	}

	return &Character{
		Server:             matches[2],
		Name:               matches[3],
		SavedVariablesPath: path,
		Times:              make(map[int]*Level),
	}
}
