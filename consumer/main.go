package main

import (
	"L0_Case/consumer/api"
	"L0_Case/consumer/inner/repo"
	"L0_Case/consumer/models"
	"context"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
)

func main() {
	var order *models.Order
	dsn := "host=localhost user=postgres password=Aa23768Aaa dbname=NATS port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := repo.GormConnect(dsn)
	if err != nil {
		log.Fatalf("err from gorm connections %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	server := new(api.Server)
	handler := new(api.Handler)
	caches := repo.New(db)
	repos := repo.NewRepository(db, caches)

	sc, _ := stan.Connect("prod", "sub-1")
	defer sc.Close()

	sc.Subscribe("static", func(m *stan.Msg) {
		message := m.Data
		err := json.Unmarshal(message, &order)
		if err != nil {
			panic(err)
		}
		//db.InsertRow(order)
		id, err := repos.Db.InsertRow(order)
		if err != nil {
			log.Fatalf("DB: %s", err)
		}
		repo.Cache.Insert(*order, id)
	})

	if err := server.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error to running server %s", err.Error())
	}

	err = caches.Upload(ctx)
	if err != nil {
		log.Fatalf("cache wasn't uploaded: %s", err)
	}

	//handler := &api.Handler{Db: db}
	//router, err := handler.Db.GetRowById(4)
	//time.Sleep(5 * time.Second)

}
