package stats

import (
	"strconv"
	"time"

	log "github.com/lobsterbandit/wow-splits/internal/logger"
	"github.com/lobsterbandit/wow-splits/pkg/aggregator"
)

func CalculateStats(characters []*aggregator.Character) {
	for _, char := range characters {
		log.Logger.Printf("Stats for %q", char.Name)
		log.Logger.Printf("Max level attained = %v", len(char.Times))

		for level := 1; level <= len(char.Times); level++ {
			time := char.Times[level]
			log.Logger.Printf("Time for level %v = %v, Aggregate = %v",
				level, toReadableLevelTime(time), toReadableAggregateTime(time))
		}
	}
}

func toReadableLevelTime(levelTime *aggregator.Level) time.Duration {
	d, _ := time.ParseDuration(strconv.Itoa(levelTime.LevelTime) + "s")
	return d
}

func toReadableAggregateTime(levelTime *aggregator.Level) time.Duration {
	d, _ := time.ParseDuration(strconv.Itoa(levelTime.AggregateTime) + "s")
	return d
}
