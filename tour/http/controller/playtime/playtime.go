package playtime

import (
	"fmt"
	"time"

	routing "github.com/qiangxue/fasthttp-routing"
)

func PlayTimeCtrl(c *routing.Context) error {

	c.SetStatusCode(200)

	t := time.Now()
	tFormat := t.Format("2006-01-02 15:04:05.0000")

	aa := fmt.Sprintf("%s %s", "CurrentTime :", tFormat)

	c.SetBodyString(aa)

	return nil
}
