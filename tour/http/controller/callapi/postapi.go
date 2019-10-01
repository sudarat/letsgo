package callapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	routing "github.com/qiangxue/fasthttp-routing"
)

func PostMyDataCtrl(c *routing.Context) error {
	input := c.Request.Body()

	responseApp, err := postData(input)
	if err != nil {
		return err
	}

	data := makeData(responseApp)

	response, err := json.Marshal(data)
	if err != nil {
		return err
	}
	c.SetStatusCode(200)
	c.SetBodyString(string(response))

	return nil
}

// func postData(input [byte]) error {
func postData(input []byte) (ResponseApp, error) {
	responseApp := ResponseApp{}

	responseRaw, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(input))
	if err != nil {
		return responseApp, err
	}

	data, err := ioutil.ReadAll(responseRaw.Body)
	err = json.Unmarshal(data, &responseApp)
	if err != nil {
		return responseApp, err
	}
	return responseApp, nil

}

func makeData(responseApp ResponseApp) Data {
	return Data{FirstName: responseApp.Json.FirstName, LastName: responseApp.Json.LastName}
}
