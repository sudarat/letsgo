package main

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	sqlCode "./tour/http/controller/sql"
)

// fetchWallet()
func TestFetchWallet(t *testing.T) {
	wantCardNumber := "1001"
	wantName := "Prare"

	db, err := sql.Open("mysql", "root:MySQL123@tcp(127.0.0.1:3306)/GoTour")
	if err != nil {
		t.Errorf("Fail: Cannot connect db")
	}
	wallet, err := sqlCode.FetchWallet(db, wantCardNumber)
	if err != nil {
		t.Errorf("Fail: FetchWallet has error")
	}

	if wallet.CardNumber != wantCardNumber {
		t.Errorf("Fail: CardNumber got %v want %v", wallet.CardNumber, wantCardNumber)
	}

	if wallet.Name != wantName {
		t.Errorf("Fail: Name got %v want %v", wallet.Name, wantName)
	}
}
