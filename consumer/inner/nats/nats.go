package nats

import (
	"context"
	"errors"
	"github.com/nats-io/stan.go"
	"sync"
)

type Streaming interface {
	GetMessage() ([]byte, error)
}

type Connector struct {
	Conn stan.Conn
}

func Connecting(ctx context.Context) (*Connector, error) {
	sc, err := stan.Connect("prod", "cl-2")
	if err != nil {
		return nil, err
	}
	return &Connector{Conn: sc}, nil
}

func (c *Connector) GetMessage() ([]byte, error) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var message []byte

	sub, err := c.Conn.Subscribe("static", func(m *stan.Msg) {
		message = m.Data
		wg.Done()
	})
	if err != nil {
		return nil, errors.New("cant receive message")
	}

	wg.Wait()
	err = sub.Unsubscribe()
	if err != nil {
		return nil, errors.New("can't unsubscribe")
	}
	return message, nil
}

func (c *Connector) Close() error {
	err := c.Conn.Close()
	if err != nil {
		return err
	}
	return nil
}
