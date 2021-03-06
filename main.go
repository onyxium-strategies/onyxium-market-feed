package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"github.com/onyxium-strategies/onyxium-strategy-worker/models"
	"log"
	"time"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("onyxium").C("market")

	for {
		time.Sleep(time.Second)
		market := make(map[string]models.Market)
		market, err = getMarketSummary()
		if err != nil {
			fmt.Println(err)
			return
		}
		// for debugging
		fmt.Println("Got market summary")

		err = c.Insert(&market)
		if err != nil {
			log.Fatal(err)
		}
	}
}
