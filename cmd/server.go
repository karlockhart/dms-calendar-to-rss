package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/karlockhart/overly-complicated-ical-parser-go/pkg/ical2"
	"gopkg.in/redis.v5"
)

func getCalendarData() {
	c, err := ical2.ParseICal2Url("https://calendar.dallasmakerspace.org/events/feed")
	if err != nil {
		log.Println("Error getting calendar.")
		return
	}

	j, err := json.Marshal(c)
	if err != nil {
		log.Println("Error converting calendar to JSON.")
	}

	json := string(j)

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	client.LPush("dms-calendar", json)
}

func main() {
	getCalendarData()
}
