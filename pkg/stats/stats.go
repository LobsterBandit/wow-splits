package stats

import (
	"github.com/lobsterbandit/wow-splits/pkg/aggregator"
	log "github.com/lobsterbandit/wow-splits/pkg/logger"
)

func CalculateStats(characters []*aggregator.Character) {
	for _, char := range characters {
		log.Logger.Printf("Stats for %q", char.Name)
		log.Logger.Printf("Max level attained = %v", len(char.Times))
	}
}
