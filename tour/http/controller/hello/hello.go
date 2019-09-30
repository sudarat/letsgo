package hello

import (
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
)

func HelloCtrl(c *routing.Context) error {
	fmt.Println("Hello PP.")

	c.SetStatusCode(200)

	response := fmt.Sprintf("%s", Hello())
	c.SetBodyString(response)

	return nil
}

func Hello() string {
	return "Hello"
}
