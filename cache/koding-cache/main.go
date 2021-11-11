package main

import (
	"log"
	"time"

	"github.com/koding/cache"
)

func main() {
	// create a cache with 2 second TTL
	mem := cache.NewMemoryWithTTL(2 * time.Second)

	// start garbage collection for expired keys
	mem.StartGC(time.Millisecond * 10)

	// set maintenance data
	err := mem.Set("maintenance_mode", "true")
	if err != nil {
		log.Println("Error when get maintenance mode data")
	}

	log.Println("The cache should be still there, let's get it..")

	// get maintenance data
	data, err := mem.Get("maintenance_mode")
	if err != nil {
		log.Println("Error when get maintenance mode data -> err : ", err)
	}
	log.Println("maintenance_mode", data)

	time.Sleep(3 * time.Second)
	log.Println("The cache should be gone, let's check it..")

	// get maintenance data
	data, err = mem.Get("maintenance_mode")
	if err != nil {
		log.Println("Error when get maintenance mode data -> err :", err)
	}
	log.Println("maintenance_mode", data)
}
