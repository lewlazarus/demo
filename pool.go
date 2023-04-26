package main

import (
	"demo/ops"
	"github.com/mailru/easyjson"
	"log"

	"github.com/nats-io/nats.go"

	"demo/data"
	m "demo/msg"
)

// Pool is an implementation of
type Pool struct {
	subject string
	size    int

	subscriptions []*nats.Subscription
}

func NewPool(subject string, size int) *Pool {
	return &Pool{
		subject:       subject,
		size:          size,
		subscriptions: make([]*nats.Subscription, 0, size),
	}
}

func (r *Pool) Run(conn *nats.Conn, storage data.StorageInterface) error {
	for i := 0; i < r.size; i++ {
		subscription, err := conn.Subscribe(r.subject, func(msg *nats.Msg) {
			r.Subscription(msg, storage)
		})
		if err != nil {
			return err
		}

		r.subscriptions = append(r.subscriptions, subscription)
	}

	return nil
}

func (r *Pool) Stop() {
	for _, subscription := range r.subscriptions {
		if err := subscription.Unsubscribe(); err != nil {
			log.Printf("nats unsubscribe error")
			log.Println(err)
		}
	}

	log.Println("pool closed")
}

func (r *Pool) Subscription(msg *nats.Msg, storage data.StorageInterface) {
	req := &m.Request{}
	if err := easyjson.Unmarshal(msg.Data, req); err != nil {
		log.Println("parsing error")
		log.Println(err)

		_ = msg.Respond([]byte("N/A"))
		return
	}

	var res float64

	switch req.Operation {
	case m.OpSum:
		res = ops.Sum(req.Values)
	case m.OpMultiply:
		res = ops.Multiply(req.Values)
	case m.OpStdDev:
		res = ops.StandardDeviation(req.Values)
	case m.OpMean:
		res = ops.Mean(req.Values)
	case m.OpVariance:
		res = ops.Variance(req.Values)
	}

	_ = storage.BeginTx()
	prev, _ := storage.Get()
	ratio := res / prev

	if res == 0 {
		_ = storage.Set(1)
	} else {
		_ = storage.Set(res)
	}

	_ = storage.EndTx()

	resp := &m.Response{
		Res:   res,
		Ratio: ratio,
	}

	d, _ := easyjson.Marshal(resp)
	_ = msg.Respond(d)
}
