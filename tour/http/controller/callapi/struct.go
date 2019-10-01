package callapi

// ResponseApp use for result of "https://httpbin.org/ip"
type ResponseApp struct {
	Origin string `json:"origin"`
	Json   Data   `json:"json,omitempty"`
}

type Data struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
}

// ResponseAPI use for result of GetIPAddressCtrl()
type ResponseAPI struct {
	IPAddress string `json:"ipAddress"`
}
