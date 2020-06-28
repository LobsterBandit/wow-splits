package account

type CollectionTotals struct {
	Accounts   int
	Servers    int
	Characters int
}

func (w *WowInstall) CalculateTotals() CollectionTotals {
	var serverTotal, characterTotal int
	for _, account := range w.Accounts {
		serverTotal += account.ServerCount
		characterTotal += account.CharacterCount
	}

	return CollectionTotals{
		Accounts:   len(w.Accounts),
		Servers:    serverTotal,
		Characters: characterTotal,
	}
}
