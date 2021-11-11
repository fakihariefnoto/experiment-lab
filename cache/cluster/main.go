package main

import (
	"log"

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
	}
	return slots, nil
}
