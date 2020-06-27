package character

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	log "github.com/lobsterbandit/wow-splits/internal/logger"
	"github.com/lobsterbandit/wow-splits/pkg/file"
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
var levelRegexp = regexp.MustCompile(`(?is).*SpeedrunSplits\s*=\s*{(.+)}.*`)
var levelDataRegexp = regexp.MustCompile(`\s*([0-9]+),\s*--\s*\[([0-9]+)\]`)

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

func (c *Character) ParseCharacterData() error {
	data, err := file.ReadFile(c.SavedVariablesPath)
	if err != nil {
		return err
	}

	levelData := levelRegexp.FindStringSubmatch(data)

	if len(levelData) <= 1 {
		return fmt.Errorf("ParseCharacter: no level data to parse")
	}

	c.ParseLevels(strings.TrimSpace(levelData[1]))

	return nil
}

func (c *Character) ParseLevels(table string) {
	var previousLevelTime int

	scanner := bufio.NewScanner(strings.NewReader(table))
	for scanner.Scan() {
		level := ParseLevelData(scanner.Text(), previousLevelTime)
		c.Times[level.Level] = level
		previousLevelTime = level.AggregateTime
	}
}

// in form:
// aggregate time, -- [level]
// 12345, 		   -- [18].
func ParseLevelData(levelText string, previous int) *Level {
	log.Logger.Printf("Parsing level text: %s", levelText)

	matches := levelDataRegexp.FindStringSubmatch(levelText)

	level, _ := strconv.Atoi(matches[2])
	time, _ := strconv.Atoi(matches[1])

	return &Level{
		Level:         level,
		LevelTime:     time - previous,
		AggregateTime: time,
	}
}
