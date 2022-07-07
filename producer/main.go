package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"time"
)

func main() {
	sc, err := stan.Connect("prod", "simple-pub")
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	value, err := ioutil.ReadFile("./producer/model.json")
	if err != nil {
		panic(err)
	}

	for i := 1; i < 101; i++ {
		err := sc.Publish("static", value)
		if err != nil {
			panic(err)
		}
		fmt.Println("Message: ", i)
		time.Sleep(3 * time.Second)
	}

}
