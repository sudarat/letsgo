package main

import (
	"fmt"
	path "./tour/routes"

	"github.com/valyala/fasthttp"

	routing "github.com/qiangxue/fasthttp-routing"
)

func main() {
	fmt.Println("Start Service.....")

	r := routing.New()
	r.NotFound(notFoundHandler)

	path.RoutePaht(r)

	h := r.HandleRequest
	h = fasthttp.CompressHandler(h)

	panic(fasthttp.ListenAndServe(":4005", h))

}

func notFoundHandler(c *routing.Context) error {
	c.RequestCtx.NotFound()
	c.Next()
	return nil
}
