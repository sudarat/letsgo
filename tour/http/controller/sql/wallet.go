package sql

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	routing "github.com/qiangxue/fasthttp-routing"
)

func GetWalletCtrl(c *routing.Context) error {
	cardNumber := c.QueryArgs().Peek("CardNumber")
	if len(string(cardNumber)) <= 0 {
		return fmt.Errorf("Require Card Number")
	}
	db, err := sql.Open("mysql", "root:MySQL123@tcp(127.0.0.1:3306)/GoTour")
	if err != nil {
		c.SetStatusCode(500)
		return err
	}
	defer db.Close()

	wallet, err := fetchWallet(db, string(cardNumber))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return fmt.Errorf("CardNumber not found")
		}
		return err
	}

	response, err := json.Marshal(wallet)
	if err != nil {
		return err
	}

	c.SetStatusCode(200)
	c.SetBodyString(string(response))
	return nil

}

func fetchWallet(db *sql.DB, cardNumber string) (Wallet, error) {
	wallet := Wallet{}
	err := db.QueryRow(strFetchWallet(), cardNumber).Scan(&wallet.CardNumber, &wallet.Name, &wallet.Age, &wallet.Money)
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}
