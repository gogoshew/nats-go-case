package main

import (
	"L0_Case/consumer/models"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"time"
)

func main() {
	var order models.Order
	sc, _ := stan.Connect("prod", "sub-1")
	defer sc.Close()

	sc.Subscribe("static", func(m *stan.Msg) {
		Data := m.Data
		err := json.Unmarshal(Data, &order)
		if err != nil {
			panic(err)
		}
		fmt.Println(order.ID,
			order.Shardkey,
			order.CustomerId)
	})
	time.Sleep(5 * time.Second)

}
