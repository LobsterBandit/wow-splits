package stats

import (
	log "github.com/lobsterbandit/wow-splits/internal/logger"
	"github.com/lobsterbandit/wow-splits/pkg/account"
)

func CalculateStats(characters []*account.Character) {
	for _, char := range characters {
		log.Logger.Printf("Stats for %q", char.Name)
		log.Logger.Printf("Max level attained = %v", len(char.Times))

		for levelNum := 1; levelNum <= len(char.Times); levelNum++ {
			level := char.Times[levelNum]
			log.Logger.Printf("Time for level %v = %v, Aggregate = %v",
				level.Level, level.ToReadableLevelTime(), level.ToReadableAggregateTime())
		}
	}
}
