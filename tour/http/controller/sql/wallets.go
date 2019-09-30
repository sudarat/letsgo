package sql

import (
	"encoding/json"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	routing "github.com/qiangxue/fasthttp-routing"
)

func GetListWalletCtrl(c *routing.Context) error {

	db, err := sql.Open("mysql", "root:MySQL123@tcp(127.0.0.1:3306)/GoTour")
	if err != nil {
		c.SetStatusCode(500)
		c.SetBodyString(err.Error())
		return err

	}
	defer db.Close()

	listWallet, err := fetchWalletAll(db)
	if err != nil {
		c.SetStatusCode(500)
		c.SetBodyString(err.Error())
		return err
	}
	jsonListWallet, err := json.Marshal(listWallet)

	c.SetStatusCode(200)
	c.SetBodyString(string(jsonListWallet))
	return nil
}

func fetchWalletAll(db *sql.DB) ([]Wallet, error) {
	list, err := db.Query(strFetchWalletAll())
	if err != nil {
		return nil, err
	}

	listWallet := []Wallet{}

	for list.Next() {
		var wallet Wallet
		err = list.Scan(&wallet.CardNumber, &wallet.Name, &wallet.Age, &wallet.Money)
		if err != nil {
			return nil, err
		}
		listWallet = append(listWallet, wallet)
	}
	defer list.Close()

	return listWallet, nil
}
