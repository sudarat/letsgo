package sql

func strFetchWalletAll() string {
	return "SELECT CardNumber, Name, IFNULL(Age,0), Money FROM GoTour.Wallet;"
}

func strFetchWallet() string {
	return "SELECT CardNumber, Name, IFNULL(Age,0), Money FROM GoTour.Wallet WHERE CardNumber = ?;"
}

func strCreateWallet() string {
	return `INSERT INTO GoTour.Wallet (CardNumber, Name, Age, Money)
	VALUE(?,?,?,?);`
}
