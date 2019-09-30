package slice

import (
	"encoding/json"
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
)

type PersonData struct {
	Name   string
	Age    int `json:"age,omitempty"`
	Weight Weight
}

type Weight float64

// MarshalJSON is Override function for json.Marshal
func (w Weight) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%0.2f", w)), nil
}

func PersonsCtrl(c *routing.Context) error {
	c.SetStatusCode(200)

	list := []PersonData{}
	list = append(list, PersonData{"Prare", 30, 60}, PersonData{"Zong", 0, 54})

	response, err := json.Marshal(list)
	if err != nil {
		return err
	}
	c.SetBodyString(string(response))

	return nil
}
