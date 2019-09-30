package sql

type Wallet struct {
	CardNumber string
	Name       string
	Age        int `json:"Age,omitempty"`
	Money      Money
}

// Money is type for money
type Money float64
