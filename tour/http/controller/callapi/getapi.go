package callapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	routing "github.com/qiangxue/fasthttp-routing"
)

// GetIPAddressCtrl is func
func GetIPAddressCtrl(c *routing.Context) error {
	responseApp, err := getIPExternal()
	if err != nil {
		return err
	}

	responseApp.Origin = strings.ReplaceAll(responseApp.Origin, " ", "")
	ipList := strings.Split(responseApp.Origin, ",")

	responseAPI := MakeIPAddressList(ipList)
	response, err := json.Marshal(responseAPI)

	if err != nil {
		return err
	}
	c.SetStatusCode(200)
	c.SetBodyString(string(response))

	return nil
}

func getIPExternal() (ResponseApp, error) {
	responseRaw, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		return ResponseApp{}, err
	}
	data, err := ioutil.ReadAll(responseRaw.Body)
	if err != nil {
		return ResponseApp{}, err
	}
	responseApp := ResponseApp{}
	err = json.Unmarshal(data, &responseApp)
	if err != nil {
		return ResponseApp{}, err
	}
	return responseApp, nil
}

func MakeIPAddressList(ipList []string) []ResponseAPI {
	responseAPI := []ResponseAPI{}
	for i := 0; i < len(ipList); i++ {
		// fmt.Println(ipList[i])
		r := ResponseAPI{IPAddress: ipList[i]}
		responseAPI = append(responseAPI, r)
	}
	return responseAPI
}
