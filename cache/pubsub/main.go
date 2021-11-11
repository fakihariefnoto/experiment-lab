package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type (
	Conn struct {
		Address  string
		Password string
	}
)

var (
	sub bool
	pub bool
)

func main() {
	conn := Conn{
		Address:  "127.0.0.1:6379",
		Password: "",
	}

	flag.BoolVar(&pub, "pub", false, "Run only publisher")
	flag.BoolVar(&sub, "sub", false, "Run only subscriber")
	flag.Parse()

	if pub {
		playPublisher(conn)
		return
	}

	if sub {
		playSubscriber(conn)
		return
	}

	log.Println("pub ", pub, "  sub ", sub)

	play(conn)

}

func connect(conn Conn) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     conn.Address,
		Password: conn.Password,
	})
	return cli
}

func play(conn Conn) {
	cli := connect(conn)

	pubsub := cli.Subscribe("mychannel")

	go func() {
		log.Println("||||||| Subscribing to channel -> ", "mychannel")
		for {
			// Wait for confirmation that subscription is created before publishing anything.
			_, err := pubsub.Receive()
			if err != nil {
				panic(err)
			}

			// Go channel which receives messages.
			ch := pubsub.Channel()
			// Consume messages.
			for msg := range ch {
				log.Println("Incoming message from :", msg.Channel, " data : ", msg.Payload)
			}
		}
	}()

	index := int(0)
	log.Println("||||||| Publishing to channel -> ", "mychannel")
	for {
		index++
		log.Println("Sending data.. ", index)
		err := cli.Publish("mychannel", "hello -> "+fmt.Sprint(index)).Err()
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 1)
	}
}

func playPublisher(conn Conn) {
	cli := connect(conn)

	log.Println("||||||| Publishing to channel -> ", "mychannel")
	index := int(0)
	for {
		index++
		log.Println("Sending data.. ", index)
		err := cli.Publish("mychannel", "hello -> "+fmt.Sprint(index)).Err()
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 1)
	}
}

func playSubscriber(conn Conn) {
	cli := connect(conn)

	pubsub := cli.Subscribe("mychannel")
	log.Println("||||||| Subscribing to channel -> ", "mychannel")
	for {
		// Wait for confirmation that subscription is created before publishing anything.
		_, err := pubsub.Receive()
		if err != nil {
			panic(err)
		}

		// Go channel which receives messages.
		ch := pubsub.Channel()
		// Consume messages.
		for msg := range ch {
			log.Println("Incoming message from :", msg.Channel, " data : ", msg.Payload)
		}
	}
}
