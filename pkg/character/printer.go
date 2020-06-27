package character

import (
	"strconv"
	"time"

	log "github.com/lobsterbandit/wow-splits/internal/logger"
)

// print discovered account/server/character structure in tree format
func PrintCharacterTree(characters []*Character, includeTimes bool) {
	accountMap := generateAccountMap(characters)

	log.Logger.Println("\n##########################################################")

	log.Logger.Printf("%d Account(s), %d Server(s) and %d Character(s) found:",
		accountMap.AccountCount, accountMap.ServerCount, accountMap.CharacterCount)

	iAccount := 1
	for accountName, serverMap := range accountMap.Accounts {
		accountSep := determineTreeSeparator(iAccount, accountMap.AccountCount)
		log.Logger.Printf("%s-- %-15s", accountSep, accountName)

		iAccount++

		iServer := 1
		for serverName, chars := range serverMap {
			serverSep := determineTreeSeparator(iServer, len(serverMap))

			if iAccount > len(serverMap) {
				accountSep = " "
			}

			log.Logger.Printf("%s%4s-- %-15s", accountSep, serverSep, serverName)

			iServer++

			for i, char := range chars {
				charSep := determineTreeSeparator(i, len(chars)-1)

				if iServer > len(serverMap) {
					serverSep = " "
				}

				if includeTimes {
					log.Logger.Printf("%s%4s%4s-- %s - %-2d%10s | %11s",
						accountSep,
						serverSep,
						charSep,
						char.Name,
						len(char.Times),
						"Level",
						"Played")

					for levelNum := 1; levelNum <= len(char.Times); levelNum++ {
						levelSep := determineTreeSeparator(levelNum, len(char.Times))

						// clear charSep if on last char
						if i == len(chars)-1 {
							charSep = " "
						}

						log.Logger.Printf("%s%4s%4s%4s-- %-2d = %11s | %11s",
							accountSep,
							serverSep,
							charSep,
							levelSep,
							char.Times[levelNum].Level,
							char.Times[levelNum].ToReadableLevelTime(),
							char.Times[levelNum].ToReadableAggregateTime())
					}
				} else {
					log.Logger.Printf("%s%4s%4s-- %s - Level %-2d",
						accountSep,
						serverSep,
						charSep,
						char.Name,
						len(char.Times))
				}
			}
		}
	}

	log.Logger.Println("##########################################################\n")
}

func (v *Level) ToReadableLevelTime() time.Duration {
	d, _ := time.ParseDuration(strconv.Itoa(v.LevelTime) + "s")
	return d
}

func (v *Level) ToReadableAggregateTime() time.Duration {
	d, _ := time.ParseDuration(strconv.Itoa(v.AggregateTime) + "s")
	return d
}

func determineTreeSeparator(n, total int) (sep string) {
	if n == total {
		sep = "\\"
	} else {
		sep = "|"
	}
	return
}

type PrintableAccountMap struct {
	Accounts       map[string]map[string][]*Character
	AccountCount   int
	ServerCount    int
	CharacterCount int
}

func generateAccountMap(characters []*Character) *PrintableAccountMap {
	accountMap := &PrintableAccountMap{
		Accounts:       make(map[string]map[string][]*Character),
		CharacterCount: len(characters),
	}
	for _, char := range characters {
		// initialize account if necessary
		if existingAccount := accountMap.Accounts[char.Account]; existingAccount == nil {
			accountMap.Accounts[char.Account] = make(map[string][]*Character)
			accountMap.AccountCount++
		}

		// initialize server slice if necessary
		if existingServer := accountMap.Accounts[char.Account][char.Server]; existingServer == nil {
			accountMap.Accounts[char.Account][char.Server] = make([]*Character, 0, 20)
			accountMap.ServerCount++
		}

		accountMap.Accounts[char.Account][char.Server] = append(accountMap.Accounts[char.Account][char.Server], char)
	}

	return accountMap
}
