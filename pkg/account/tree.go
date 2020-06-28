package account

import (
	"strconv"
	"time"

	log "github.com/lobsterbandit/wow-splits/internal/logger"
)

// print discovered account/server/character structure in tree format
func PrintAccountTree(accounts []*Account, includeTimes bool) {
	log.Logger.Println("##########################################################")

	var serverTotal, characterTotal int
	for _, account := range accounts {
		serverTotal += account.ServerCount
		characterTotal += account.CharacterCount
	}

	log.Logger.Printf("%d Account(s), %d Server(s) and %d Character(s) found:",
		len(accounts), serverTotal, characterTotal)

	for iAccount, account := range accounts {
		accountSep := determineTreeSeparator(iAccount == len(accounts)-1)
		log.Logger.Printf("%s-- %-15s", accountSep, account.Name)

		iServer := 1
		for serverName, characters := range account.Servers {
			serverSep := determineTreeSeparator(iServer == len(account.Servers))

			if iAccount == len(accounts)-1 {
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
