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

	for i := 0; i <= 100; i++ {
		err := sc.Publish("static", value)
		if err != nil {
			panic(err)
		}
		fmt.Println("Push ", i)
		time.Sleep(3 * time.Second)
	}

}
