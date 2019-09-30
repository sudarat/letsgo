package sql

import "fmt"

// MarshalJSON is override func.
func (m Money) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%0.2f", m)), nil
}
