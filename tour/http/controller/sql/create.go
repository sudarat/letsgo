package sql

import (
	"database/sql"
	"encoding/json"
	_ "fmt"

	_ "github.com/go-sql-driver/mysql"

	routing "github.com/qiangxue/fasthttp-routing"
)

func CreateWalletCtrl(c *routing.Context) error {
	body := c.Request.Body()

	var wallet Wallet
	err := json.Unmarshal(body, &wallet)
	if err != nil {
		return err
	}

	db, err := sql.Open("mysql", "root:MySQL123@tcp(127.0.0.1:3306)/GoTour")
	if err != nil {
		return err
	}
	defer db.Close()

	err = createWallet(db, wallet)
	if err != nil {
		return err
	}

	c.SetStatusCode(200)
	c.SetBodyString("{\"Result\":\"Success\"}")

	return nil
}

func createWallet(db *sql.DB, wallet Wallet) error {
	insert, err := db.Query(strCreateWallet(), wallet.CardNumber, wallet.Name, wallet.Age, wallet.Money)
	if err != nil {
		return err
	}
	defer insert.Close()
	return nil
}
