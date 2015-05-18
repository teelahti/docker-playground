package main

import (
	"os"
	"fmt"
	"gopkg.in/redis.v3"
)

var client *redis.Client

func main() {
    fmt.Println("hello world")
	
	addr := os.Getenv("DB_PORT_6379_TCP_ADDR") + ":" + os.Getenv("DB_PORT_6379_TCP_PORT")
	
	fmt.Println("FOO:", addr)
	
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}