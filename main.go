package main

import (
	"github.com/dmcleish91/go_echo_api/src/router"
)

func main() {
	e := router.New()
	e.Logger.Fatal(e.Start(":1323"))
}
