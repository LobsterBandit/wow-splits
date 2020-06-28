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

	log.Logger.Println("Collecting character leveling times...")

	accounts, err := account.CreateAllAccounts(*wowDir, *debug)
	if err != nil {
		log.Logger.Printf("Error finding accounts in %q: %v", *wowDir, err)
	}

	for _, account := range accounts {
		err := account.PopulateAccountData(*debug)

		if err != nil {
			log.Logger.Printf("Error populating account data: %v", err)
		}
	}

	if *printTree || *printTreeWithTimes {
		account.PrintAccountTree(accounts, *printTreeWithTimes)
	}
}
