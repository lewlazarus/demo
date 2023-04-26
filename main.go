package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"

	"demo/data"
)

//go:generate go run ./gen/generator.go Sum 10
//go:generate go run ./gen/generator.go Multiply 10
//go:generate go run ./gen/generator.go Mean
//go:generate go run ./gen/generator.go Variance 10
//go:generate go run ./gen/generator.go StdDev
//go:generate easyjson -all ./msg/request.go
//go:generate easyjson -all ./msg/response.go

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("can not read the .env file")
	}
}

func main() {
	conf := NewConfig()
	if err := conf.Read(); err != nil {
		log.Println("config read error")
		log.Panic(err)
	}

	storage := data.NewFileStorage()
	if err := storage.Init(); err != nil {
		log.Println("storage init error")
		log.Panic(err)
	}

	val, err := storage.Get()
	if err != nil {
		log.Println("storage init get error")
		log.Panic(err)
	}
	if val == 0 {
		_ = storage.Set(1)
	}

	conn, err := nats.Connect(conf.NutsUrl)
	if err != nil {
		log.Println("nats connection error")
		log.Panic(err)
	}

	defer conn.Close()

	pool := NewPool(conf.NutsSubject, conf.PoolSize)
	if err := pool.Run(conn, storage); err != nil {
		log.Println("pool init error")
		log.Panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		<-ch

		pool.Stop()
	}()

	wg.Wait()

	if err := storage.Persist(); err != nil {
		log.Println("storage data persist error")
		log.Println(err)
	}

	log.Println("bye!")
}
