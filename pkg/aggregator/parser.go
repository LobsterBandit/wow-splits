package aggregator

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	log "github.com/lobsterbandit/wow-splits/internal/logger"
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

func CreateCharacter(path string) *Character {
	r := regexp.MustCompile(`_classic_/WTF/Account/(?P<AccountName>\w+)/(?P<ServerName>\w+)/(?P<CharacterName>\w+)`)

	matches := r.FindStringSubmatch(path)

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
	data, err := readSpeedrunSplits(c.SavedVariablesPath)
	if err != nil {
		return err
	}

	levelRegexp := regexp.MustCompile(`(?is).*SpeedrunSplits = {(.+)}.*`)
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
		level := parseLevelData(scanner.Text(), previousLevelTime)
		c.Times[level.Level] = level
		previousLevelTime = level.AggregateTime
	}
}

// in form:
// aggregate time, -- [level]
// 12345, 		   -- [18].
func parseLevelData(levelText string, previous int) *Level {
	log.Logger.Printf("Parsing level text: %s", levelText)

	r := regexp.MustCompile(`\s*([0-9]+), -- \[([0-9]+)\]`)
	matches := r.FindStringSubmatch(levelText)

	level, _ := strconv.Atoi(matches[2])
	time, _ := strconv.Atoi(matches[1])

	return &Level{
		Level:         level,
		LevelTime:     time - previous,
		AggregateTime: time,
	}
}

// Global:
// - SpeedrunSplitsGold
// - SpeedrunSplitsPB
// - SpeedrunSplitsOptions, ignore
// Character:
// - SpeedrunSplits

func readSpeedrunSplits(path string) (data string, err error) {
	log.Logger.Printf("Reading %q", path)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Logger.Printf("Error reading %q: %v", path, err)
		return "", err
	}

	return string(content), nil
}
