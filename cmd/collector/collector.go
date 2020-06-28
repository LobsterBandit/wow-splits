package main

import (
	"flag"
	"os"

	log "github.com/lobsterbandit/wow-splits/internal/logger"
	"github.com/lobsterbandit/wow-splits/pkg/account"
)

func main() {
	wowDir := flag.String("wowdir", "", "(Required) path to \"World of Warcraft\" install `directory`")
	printTree := flag.Bool("tree", false, "print Account/Server/Character tree")
	printTreeWithTimes := flag.Bool("tree-with-times", false, "print Account/Server/Character tree with per level times")
	debug := flag.Bool("debug", false, "print debug logs to stdout")

	flag.Parse()

	log.CreateGlobalLogger(*debug)

	if *wowDir == "" {
		flag.Usage()
		os.Exit(1)
	}

	log.Logger.Println("Collecting leveling times...")

	wowInstall := &account.WowInstall{
		RootDir:  *wowDir,
		Accounts: make([]*account.Account, 0, 5),
	}

	err := wowInstall.CreateAllAccounts(*debug)
	if err != nil {
		log.Logger.Printf("Error finding accounts in %q: %v", *wowDir, err)
	}

	for _, account := range wowInstall.Accounts {
		err := account.PopulateAccountData(*debug)

		if err != nil {
			log.Logger.Printf("Error populating account data: %v", err)
		}
	}

	if *printTree || *printTreeWithTimes {
		wowInstall.PrintAccountTree(*printTreeWithTimes)
	}

	log.Logger.Println("Finished collecting leveling data")

	results := wowInstall.CalculateTotals()

	log.Logger.Printf("Found %d Account(s), %d Server(s) and %d Character(s)",
		results.Accounts, results.Servers, results.Characters)
}
