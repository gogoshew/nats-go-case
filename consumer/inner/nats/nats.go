package nats

import (
	"context"
	"errors"
	"github.com/nats-io/stan.go"
	"sync"
)

// Интерфейс для стриминга

type Streaming interface {
	GetMessage() ([]byte, error)
}

// Структура подключения

type Connector struct {
	Conn stan.Conn
}

// Функция для подключения

func Connecting(ctx context.Context) (*Connector, error) {
	sc, err := stan.Connect("prod", "client-2")
	if err != nil {
		return nil, err
	}
	return &Connector{Conn: sc}, nil
}

// Функция для получения сообщения

func (c *Connector) GetMessage() ([]byte, error) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var message []byte

	sub, err := c.Conn.Subscribe("static", func(m *stan.Msg) {
		message = m.Data
		wg.Done()
	})
	if err != nil {
		return nil, errors.New("can't receive message")
	}
	wg.Wait()
	err = sub.Unsubscribe()
	if err != nil {
		return nil, errors.New("can't unsubscribe")
	}
	return message, nil
}

// Функция закрывающая соединение

func (c *Connector) Close() error {
	err := c.Conn.Close()
	if err != nil {
		return err
	}
	return nil
}
