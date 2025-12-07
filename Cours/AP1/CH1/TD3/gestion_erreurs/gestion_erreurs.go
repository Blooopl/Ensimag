package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func main() {
	testDiv(3, 4)
	testDiv(3, 0)
}

func division(x, y float64) (float64, error) {
	var res float64
	var err error
	if y == 0 {
		res = 0
		err = errors.New("Division par 0")
	} else {
		err = nil
		res = x / y
	}
	return res, err
}

func testDiv(x, y float64) {
	fmt.Println(strconv.FormatFloat(x, 'g', -1, 64), "par", strconv.FormatFloat(y, 'g', -1, 64))
	res, err := division(x, y)
	if err == nil {
		fmt.Println(res)
	} else {
		log.SetFlags(log.Flags() | log.Lshortfile)
		log.Fatal(err)
	}
}
