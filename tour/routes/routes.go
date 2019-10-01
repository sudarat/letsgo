package routes

import (
	// "PP/tour/http/controller/fibonacci"
	// "PP/tour/http/controller/hello"
	// "PP/tour/http/controller/playtime"
	// "PP/tour/http/controller/slice"

	"../http/controller/callapi"
	"../http/controller/fibonacci"
	"../http/controller/hello"
	"../http/controller/playtime"
	"../http/controller/slice"
	"../http/controller/sql"

	routing "github.com/qiangxue/fasthttp-routing"
)

// RoutePaht for set path of app
func RoutePaht(r *routing.Router) {
	p := r.Group("/tour")
	p.Get("/hello", hello.HelloCtrl)
	p.Get("/fibonacci", fibonacci.FibonacciCtrl)
	p.Get("/timenow", playtime.PlayTimeCtrl)

	p.Get("/persons", slice.PersonsCtrl)
	p.Get("/wallets", sql.GetListWalletCtrl)
	p.Get("/wallet", sql.GetWalletCtrl)
	p.Post("/wallet", sql.CreateWalletCtrl)

	g := p.Group("/callapi")
	g.Get("/ipaddress", callapi.GetIPAddressCtrl)
	g.Post("/mydata", callapi.PostMyDataCtrl)

}
