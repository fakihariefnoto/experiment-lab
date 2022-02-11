package main

import (
	"log"
	"time"

	goredis "github.com/go-redis/redis"
	redigo "github.com/gomodule/redigo/redis"
)

type (
	Conn struct {
		Address  string
		Password string
	}
)

func main() {
	conn := Conn{
		Address:  "127.0.0.1:6379",
		Password: "",
	}

	// pipeline := false
	// playGoRedis(conn, pipeline)
	playRedigo(conn)

}

/*
	START of GOREDIS
*/

func connectGoRedis(conn Conn) *goredis.Client {
	cli := goredis.NewClient(&goredis.Options{
		Addr:     conn.Address,
		Password: conn.Password,
	})
	return cli
}

func playGoRedis(conn Conn, pipeline bool) {
	cli := connectGoRedis(conn)
	if cli == nil {
		log.Fatalln("Redis connection error")
	}

	defaultDuration := time.Duration(time.Minute * 15)

	sleep := func(msecond int) {
		time.Sleep(time.Duration(msecond * 100 * int(time.Millisecond)))
	}

	//////////// String & Integer https://redis.io/commands#string
	log.Println("\n\nSTRING TYE DATA COMMAND\n\n")

	// SET productview:1 1
	log.Println("set product view data with key : productview:1 to 12")
	cli.Set("productview:1", 12, defaultDuration)
	sleep(10)

	// GET productview:1
	res, _ := cli.Get("productview:1").Result()
	log.Println("get productview:1 ->", res)
	time.Sleep(10)

	// INCRBY productview:1 10
	log.Println("increase productview:1 -> add 1009")
	cli.IncrBy("productview:1", 1009)
	sleep(10)

	// GET productview:1
	res, _ = cli.Get("productview:1").Result()
	log.Println("get latest counter of productview:1 ->", res)
	sleep(30)

	//////////// HASHES https://redis.io/commands#hash

	log.Println("\n\nHASH TYPE DATA COMMAND\n\n")

	// HMSET (Multi set) -> productdata:1
	log.Println("Setup product data for product id 1 with key : productdata:1\n")
	mapData := map[string]interface{}{
		"name":   "the best product in tokped",
		"desc":   "buy it!",
		"price":  99999,
		"stock":  3,
		"shopID": 55,
	}
	cli.HMSet("productdata:1", mapData)
	sleep(10)

	// HGETALL
	log.Println("Get all data for productdata:1")
	mapResult := cli.HGetAll("productdata:1")
	mapres, _ := mapResult.Result()
	for key, val := range mapres {
		log.Printf("key : %v, value : %v", key, val)
	}
	sleep(20)
	log.Println("")

	// HSET
	log.Println("Update some field for name and stock for productdata:1\n")
	cli.HSet("productdata:1", "name", "the best product in tokped, number uno!!!!")
	cli.HIncrBy("productdata:1", "stock", -2)
	sleep(10)

	// HGETALL
	log.Println("Get all data for productdata:1\n")
	mapResult = cli.HGetAll("productdata:1")
	mapres = map[string]string{}
	mapres, _ = mapResult.Result()
	for key, val := range mapres {
		log.Printf("key : %v, value : %v", key, val)
	}
	sleep(10)

	// HGET
	log.Println("")
	log.Println("Get only description field\n")
	resHGET, _ := cli.HGet("productdata:1", "desc").Result()
	log.Println("description : ", resHGET)
	sleep(30)

	/////////// LIST https://redis.io/commands#list

	log.Println("\n\nLIST TYPE DATA COMMAND\n\n")

	// LIST -> LLL LL L R RR RR

	// LPUSH -> lastestproduct product1 product5 product6 product7 product8 product10
	/*
		Result
		product10
		product8
		product7
		product6
		product5
		product1
	*/
	cli.Del("lastestproduct")
	log.Println("Create list with command LPUSH and key 'lastestproduct' and value are -> product1 product5 product6 product7 product8 product10\n")
	cli.LPush("lastestproduct", "product1", "product5", "product6", "product7", "product8", "product10")
	sleep(10)

	// Showing the data
	log.Println("Get all list data on 'lastestproduct'")
	sliceLatest, _ := cli.LRange("lastestproduct", 0, -1).Result()
	for i, val := range sliceLatest {
		log.Printf("%v : %v\n", i, val)
	}
	log.Println("")
	sleep(20)

	// LPOP (remove first or left)
	log.Println("Command LPop will remove data in first index. we will remove 2 data  : product8 & product10 ")
	cli.LPop("lastestproduct")
	cli.LPop("lastestproduct")
	sleep(10)
	// Showing the data
	log.Println("Get all list data on 'lastestproduct'")
	sliceLatest, _ = cli.LRange("lastestproduct", 0, -1).Result()
	for i, val := range sliceLatest {
		log.Printf("%v : %v\n", i, val)
	}
	log.Println("")
	sleep(20)

	// RPUSH -> bestproduct product1 product5 product6 product7 product8 product10
	/*
		RESULT
		product1
		product5
		product6
		product7
		product8
		product10
	*/
	cli.Del("bestproduct")
	log.Println("Create list with command RPUSH and key 'bestproduct' and value are -> product1 product5 product6 product7 product8 product10\n")
	cli.RPush("bestproduct", "product1", "product5", "product6", "product7", "product8", "product10")
	sleep(10)

	// Showing the data
	log.Println("Get all list data on 'bestproduct'")
	sliceBest, _ := cli.LRange("bestproduct", 0, -1).Result()
	for i, val := range sliceBest {
		log.Printf("%v : %v\n", i, val)
	}
	log.Println("")
	sleep(20)

	// LPOP (remove first or left)
	log.Println("Command RPop to remove data in last index. we will remove 2 data : product8 & product10 ")
	cli.RPop("bestproduct")
	cli.RPop("bestproduct")
	sleep(10)
	// Showing the data
	log.Println("Get all list data on 'bestproduct'")
	sliceBest, _ = cli.LRange("bestproduct", 0, -1).Result()
	for i, val := range sliceBest {
		log.Printf("%v : %v\n", i, val)
	}
	log.Println("")
	sleep(20)

	////////  PIPELINE
	log.Println("REDIS PIPELINE")
	log.Println("Lets try redis pipeline to reduce network call and latency, lets compare the time")
	log.Println("Lets use exisiting key, productview:1 and productdata:1")
	log.Println("update productview:1 by 50000 (with multi call) and productdata:1 field name to `updated w/ and w/out pipeline' w/ and w/out pipeline")
	startTime := time.Now()
	pipe := cli.Pipeline()

	pipe.IncrBy("productview:1", 5000)
	pipe.IncrBy("productview:1", 5000)
	pipe.IncrBy("productview:1", 5000)
	pipe.IncrBy("productview:1", 10000)
	pipe.HSet("productdata:1", "name", "updated w/ pipeline")

	pipe.Exec()
	log.Println("Pipeline has been executed! with ", time.Since(startTime))

	startTime = time.Now()
	cli.IncrBy("productview:1", 5000)
	cli.IncrBy("productview:1", 5000)
	cli.IncrBy("productview:1", 5000)
	cli.IncrBy("productview:1", 10000)
	cli.HSet("productdata:1", "name", "updated w/out pipeline")
	log.Println("Updated without pipeline! with ", time.Since(startTime))

	// GET productview:1
	res, _ = cli.Get("productview:1").Result()
	log.Println("get latest counter of productview:1 ->", res)
	sleep(30)

	log.Println("Get only name field")
	resHGET, _ = cli.HGet("productdata:1", "name").Result()
	log.Println("name : ", resHGET)
	sleep(30)
}

/*
	END of GOREDIS
*/

/*
	START of REDIGO
*/

func connectRedigo(conn Conn) redigo.Conn {
	c, _ := redigo.Dial("tcp", conn.Address)
	return c
}

func playRedigo(conn Conn) {
	cli := connectRedigo(conn)

	log.Println("Set new key productview:2 to 6")
	cli.Do("SET", "productview:2", 6)
	log.Println("Increase productview:2 by 25")
	cli.Do("INCRBY", "productview:2", 25)

	log.Println("Get exisiting productview:1 data")
	data, _ := redigo.Bytes(cli.Do("GET", "productview:1"))
	log.Println("Data for productview:1 is ", string(data))
	data, _ = redigo.Bytes(cli.Do("GET", "productview:2"))
	log.Println("Data for productview:2 is ", string(data))

	////////  PIPELINE
	log.Println("REDIS PIPELINE")
	log.Println("Lets try redis pipeline to reduce network call and latency, lets compare the time")
	log.Println("Lets use exisiting key, productview:2")
	log.Println("update productview:2 by 50000 (with multi call) w/ and w/out pipeline")
	startTime := time.Now()
	log.Println("Pipeline started")

	cli.Send("INCRBY", "productview:2", 10000)
	cli.Send("INCRBY", "productview:2", 10000)
	cli.Send("INCRBY", "productview:2", 5000)
	cli.Flush()

	log.Println("Pipeline done! with ", time.Since(startTime))

	startTime = time.Now()
	cli.Do("INCRBY", "productview:2", 10000)
	cli.Do("INCRBY", "productview:2", 10000)
	cli.Do("INCRBY", "productview:2", 5000)
	log.Println("Updated without pipeline! with ", time.Since(startTime))
	data, _ = redigo.Bytes(cli.Do("GET", "productview:2"))
	log.Println("Data for productview:2 is ", string(data))

}

/*
	END oF REDIGO
*/
