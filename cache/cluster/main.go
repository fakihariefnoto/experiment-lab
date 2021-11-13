package main

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

func main() {

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		ClusterSlots:  clusterSlots,
		RouteRandomly: true,
	})

	err := rdb.Ping()

	if err != nil {
		log.Println(err)
	}

	defaultDuration := time.Duration(time.Minute * 15)

	sleep := func(msecond int) {
		time.Sleep(time.Duration(msecond * 100 * int(time.Millisecond)))
	}

	//////////// String & Integer https://redis.io/commands#string
	log.Println("\n\nSTRING TYE DATA COMMAND\n\n")

	// SET productview:1 1
	log.Println("set product view data with key : productview:1 productview:2 productview:4")
	rdb.Set("productview:1", 12, defaultDuration)
	rdb.Set("productview:2", 12, defaultDuration)
	rdb.Set("productview:4", 12, defaultDuration)
	sleep(10)

	// GET productview:1
	res, _ := rdb.Get("productview:1").Result()
	log.Println("get productview:1 ->", res)
	res, _ = rdb.Get("productview:2").Result()
	log.Println("get productview:2 ->", res)
	res, _ = rdb.Get("productview:4").Result()
	log.Println("get productview:4 ->", res)
	time.Sleep(10)

}

// maximum slots : 16384
func clusterSlots() ([]redis.ClusterSlot, error) {
	slots := []redis.ClusterSlot{
		// First node with 1 master and 1 slave.
		{
			Start: 0,
			End:   8191,
			Nodes: []redis.ClusterNode{{
				Addr: ":7000", // master
			}, {
				Addr: ":7001", // 1st slave
			}},
		},
		// Second node with 1 master and 1 slave.
		{
			Start: 8192,
			End:   16383,
			Nodes: []redis.ClusterNode{{
				Addr: ":8000", // master
			}, {
				Addr: ":8001", // 1st slave
			}},
		},
		// Second node with 1 master and 1 slave.
		{
			Start: 8192,
			End:   16383,
			Nodes: []redis.ClusterNode{{
				Addr: ":9000", // master
			}, {
				Addr: ":9001", // 1st slave
			}},
		},
	}
	return slots, nil
}
