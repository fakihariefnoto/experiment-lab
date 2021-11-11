package main

import (
	"errors"
	"log"
	"time"
)

type (
	cache struct {
		data   map[string]data
		ttl    int
		init   bool
		create bool
	}

	data struct {
		d interface{}
		t time.Time
	}
)

func main() {
	// create cache system with ttl 2 second
	c := CreateInMemCache(2)

	// initiate GC with 10 milisecond run
	c.InitGC(10)

	// set data integer
	c.Set("Tes-data-int", 1)
	// set data string
	c.Set("Tes-data-string", "wow")

	d, err := c.Get("Tes-data-int")
	if err != nil {
		log.Println("Error -> ", err)
		return
	}
	log.Println("Data int -> ", d)

	ds, err := c.Get("Tes-data-string")
	if err != nil {
		log.Println("Error -> ", err)
		return
	}
	log.Println("Data string -> ", ds)

	log.Println("Look after 3 second, to see if gc works")
	time.Sleep(3 * time.Second)

	d, err = c.Get("Tes-data-int")
	if err != nil {
		log.Println("Error -> ", err)
	}
	log.Println("Data int -> ", d)

	ds, err = c.Get("Tes-data-struct")
	if err != nil {
		log.Println("Error -> ", err)
	}

	log.Println("Data struct -> ", ds)

}

// CreateInMemCache to create cache object, duration should be in second
func CreateInMemCache(second int) *cache {
	return &cache{
		data:   make(map[string]data),
		ttl:    second,
		create: true,
	}

}

// InitGC to initiate GC cycle, duration should be in milisecond
func (c *cache) InitGC(d int) {
	c.init = true
	go func() {
		for {
			time.Sleep(time.Millisecond * time.Duration(d))
			go func() {
				c.startGC()
			}()
		}
	}()
}

func (c *cache) startGC() {
	now := time.Now()
	for key, value := range c.data {
		if now.After(value.t.Add(time.Second * time.Duration(c.ttl))) {
			delete(c.data, key)
		}
	}
}

func (c *cache) Set(key string, value interface{}) error {
	if !c.init && !c.create {
		return errors.New("error init first")
	}
	c.data[key] = data{
		d: value,
		t: time.Now(),
	}
	return nil
}

func (c *cache) Get(key string) (interface{}, error) {
	if !c.init && !c.create {
		return nil, errors.New("error init first")
	}

	if data, val := c.data[key]; val {
		now := time.Now()
		if now.After(data.t.Add(time.Second * time.Duration(c.ttl))) {
			return nil, errors.New("expired")
		}
		return data.d, nil
	}

	return nil, errors.New("not found")
}
