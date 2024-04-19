package main

import (
	hellodemo "github.com/Printemps417/optionloader/examples/hello/kitex_gen/hellodemo/hello"
	"log"
)

func main() {
	svr := hellodemo.NewServer(new(HelloImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
