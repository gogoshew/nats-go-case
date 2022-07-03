package main

import (
	"L0_Case/consumer/inner/repo"
	"L0_Case/consumer/models"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

func main() {
	var order models.Order
	dsn := "host=localhost user=iabalymov password=Aa23768Aaa dbname=NATS port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := repo.GormConnect(dsn)
	if err != nil {
		log.Fatalf("err from gorm connections %s", err)
	}

	sc, _ := stan.Connect("prod", "sub-1")
	defer sc.Close()

	sc.Subscribe("static", func(m *stan.Msg) {
		message := m.Data
		err := json.Unmarshal(message, &order)
		if err != nil {
			panic(err)
		}

	})

	time.Sleep(5 * time.Second)

}
