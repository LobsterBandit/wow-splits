package account

import (
	"strconv"
	"time"

	log "github.com/lobsterbandit/wow-splits/internal/logger"
)

// print discovered account/server/character structure in tree format
func (w *WowInstall) PrintAccountTree(includeTimes bool) {
	log.Logger.Println("##########################################################")

	totals := w.CalculateTotals()

	log.Logger.Printf("%d Account(s), %d Server(s) and %d Character(s):",
		totals.Accounts, totals.Servers, totals.Characters)

	for iAccount, account := range w.Accounts {
		isLastAccount := iAccount == totals.Accounts-1
		accountSep := determineTreeSeparator(isLastAccount)
		log.Logger.Printf("%s-- %-15s", accountSep, account.Name)

		iServer := 1
		for serverName, characters := range account.Servers {
			serverSep := determineTreeSeparator(iServer == len(account.Servers))

			if isLastAccount {
				accountSep = " "
			}

			log.Logger.Printf("%s%4s-- %-15s", accountSep, serverSep, serverName)

			iServer++

			for iCharacter, character := range characters {
				charSep := determineTreeSeparator(iCharacter == len(characters)-1)

				if iServer > len(account.Servers) {
					serverSep = " "
				}

				if includeTimes {
					log.Logger.Printf("%s%4s%4s-- %s - %-2d%10s | %11s",
						accountSep,
						serverSep,
						charSep,
						character.Name,
						len(character.Times),
						"Level",
						"Played")

					for levelNum := 1; levelNum <= len(character.Times); levelNum++ {
						levelSep := determineTreeSeparator(levelNum == len(character.Times))

						// clear charSep if on last char
						if iCharacter == len(characters)-1 {
							charSep = " "
						}

						log.Logger.Printf("%s%4s%4s%4s-- %-2d = %11s | %11s",
							accountSep,
							serverSep,
							charSep,
							levelSep,
							character.Times[levelNum].Level,
							character.Times[levelNum].ToReadableLevelTime(),
							character.Times[levelNum].ToReadableAggregateTime())
					}
				} else {
					log.Logger.Printf("%s%4s%4s-- %s - Level %-2d",
						accountSep,
						serverSep,
						charSep,
						character.Name,
						len(character.Times))
				}
			}
		}
	}

	log.Logger.Println("##########################################################")
}

func (v *Level) ToReadableLevelTime() time.Duration {
	d, _ := time.ParseDuration(strconv.Itoa(v.LevelTime) + "s")
	return d
}

func (v *Level) ToReadableAggregateTime() time.Duration {
	d, _ := time.ParseDuration(strconv.Itoa(v.AggregateTime) + "s")
	return d
}

func determineTreeSeparator(isLast bool) (sep string) {
	if isLast {
		sep = "+"
	} else {
		sep = "|"
	}
	return
}
