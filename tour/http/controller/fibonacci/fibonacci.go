package fibonacci

import (
	"errors"
	"fmt"
	"strconv"

	routing "github.com/qiangxue/fasthttp-routing"
)

// FibonacciCtrl for control request
func FibonacciCtrl(c *routing.Context) error {
	c.SetStatusCode(200)

	start := c.QueryArgs().Peek("start")
	end := c.QueryArgs().Peek("end")

	fmt.Println("Start : ", start)

	dataS := string(start)
	if dataS == "" {
		err := errors.New("Start is require")
		return err
	}

	dataE := string(end)

	startI, _ := strconv.Atoi(dataS)
	endI, _ := strconv.Atoi(dataE)

	if endI <= 0 {
		endI = startI + 10
	}

	response := Fibonacci(startI, endI)
	c.SetBodyString(response)
	return nil
}

// Fibonacci for control flow
func Fibonacci(start int, end int) string {
	fibo := fibonacciCal()

	result := ""
	for i := start; i < end; i++ {
		result = result + " " + strconv.Itoa(fibo(i))
	}

	return result
}

// fibonacciCal for cal n
func fibonacciCal() func(int) int {
	return func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		return fibonacciCal()(n-1) + fibonacciCal()(n-2)
	}
}
