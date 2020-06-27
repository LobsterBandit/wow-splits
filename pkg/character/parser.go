package character

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/lobsterbandit/wow-splits/pkg/file"
)

var levelRegexp = regexp.MustCompile(`([0-9]+),\s*--\s*\[{1}([0-9]+)\]{1}`)

func (c *Character) ParseCharacterData(debug bool) error {
	data, err := file.ReadFile(c.SavedVariablesPath, debug)
	if err != nil {
		return err
	}

	// levelData = [[ 0 = full match, 1 = aggregate time, 2 = level],...]
	levelData := levelRegexp.FindAllStringSubmatch(data, -1)

	if levelData == nil {
		return fmt.Errorf("ParseCharacter: no level data to parse")
	}

	c.ParseLevelMatches(levelData)

	return nil
}

func (c *Character) ParseLevelMatches(levels [][]string) {
	var previousLevelTime int

	for _, match := range levels {
		aggregateTime, _ := strconv.Atoi(match[1])
		level, _ := strconv.Atoi(match[2])

		c.Times[level] = &Level{
			Level:         level,
			LevelTime:     aggregateTime - previousLevelTime,
			AggregateTime: aggregateTime,
		}

		previousLevelTime = aggregateTime
	}
}
