package aggregator

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	log "github.com/lobsterbandit/wow-splits/pkg/logger"
)

type Level struct {
	Level         int
	LevelTime     int
	AggregateTime int
}

type Character struct {
	Account string
	Server  string
	Name    string
	Class   string
	Times   map[int]*Level
}

// Global:
// - SpeedrunSplitsGold
// - SpeedrunSplitsPB
// - SpeedrunSplitsOptions, ignore
// Character:
// - SpeedrunSplits

func ReadSpeedrunSplits(path string) (data string, err error) {
	log.Logger.Printf("Reading file at %q", path)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Logger.Printf("Error reading %q: %v", path, err)
		return "", err
	}

	return string(content), nil
}

func ParseCharacter(path string) (*Character, error) {
	character := parseCharacterMetadata(path)
	if character == nil {
		return &Character{}, fmt.Errorf("ParseCharacter: cannot parse account/server/character from path")
	}

	data, err := ReadSpeedrunSplits(path)
	if err != nil {
		log.Logger.Printf("\nerror reading %q: %v", path, err)
		return character, err
	}

	levelRegexp := regexp.MustCompile(`(?is).*SpeedrunSplits = {(.+)}.*`)
	levelData := levelRegexp.FindStringSubmatch(data)

	if len(levelData) <= 1 {
		return character, nil
	}

	return parseCharacterLevels(character, strings.TrimSpace(levelData[1])), nil
}

func parseCharacterMetadata(path string) *Character {
	r := regexp.MustCompile(`_classic_/WTF/Account/(?P<AccountName>\w+)/(?P<ServerName>\w+)/(?P<CharacterName>\w+)`)

	matches := r.FindStringSubmatch(path)

	if matches == nil {
		return nil
	}

	return &Character{
		Account: matches[1],
		Server:  matches[2],
		Name:    matches[3],
		Times:   make(map[int]*Level),
	}
}

func parseCharacterLevels(character *Character, table string) *Character {
	var aggregateTime int

	scanner := bufio.NewScanner(strings.NewReader(table))
	for scanner.Scan() {
		level := parseLevel(scanner.Text(), aggregateTime)
		character.Times[level.Level] = level
		aggregateTime += level.LevelTime
	}

	return character
}

// in form:
// time for level, -- [level]
// 12345, -- [18].
func parseLevel(levelText string, aggregateToLevel int) *Level {
	r := regexp.MustCompile(`\s*([0-9]+), -- \[([0-9]+)\]`)
	matches := r.FindStringSubmatch(levelText)

	level, _ := strconv.Atoi(matches[2])
	time, _ := strconv.Atoi(matches[1])

	return &Level{
		Level:         level,
		LevelTime:     time,
		AggregateTime: aggregateToLevel + time,
	}
}
