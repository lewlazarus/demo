package main

import (
	"log"

	"github.com/mailru/easyjson"
	"github.com/nats-io/nats.go"

	m "demo/msg"
	"demo/process"
)

// Pool is an implementation of
type Pool struct {
	subject string
	size    int

	processor     *process.Processor
	subscriptions []*nats.Subscription
}

func NewPool(subject string, size int, processor *process.Processor) *Pool {
	return &Pool{
		subject:       subject,
		size:          size,
		processor:     processor,
		subscriptions: make([]*nats.Subscription, 0, size),
	}
}

func (r *Pool) Run(conn *nats.Conn) error {
	for i := 0; i < r.size; i++ {
		subscription, err := conn.Subscribe(r.subject, func(msg *nats.Msg) {
			r.Subscription(msg)
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

func (r *Pool) Subscription(msg *nats.Msg) {
	req := &m.Request{}
	if err := easyjson.Unmarshal(msg.Data, req); err != nil {
		log.Println("parsing error occurred")
		log.Println(err)

		_ = msg.Respond([]byte("N/A"))
		return
	}

	res, ratio, err := r.processor.Process(req)
	if err != nil {
		log.Println("processing error occurred")
		log.Println(err)

		_ = msg.Respond([]byte("N/A"))
		return
	}

	d, _ := easyjson.Marshal(m.NewResponse(res, ratio))
	_ = msg.Respond(d)
}
