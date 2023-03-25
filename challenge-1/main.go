package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"secure/challenge-1/domain"
	"secure/challenge-1/routers"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init router
	router := gin.Default()

	// Elements router
	routers.ElementRouters(router)
	
	// Post simulation
	go sumulation()

	// Router port
	err := router.Run(":4000")

	// Handle router
	if err != nil {
		panic("Error When Running")
	}
}

func sumulation(){
	for {
		time.Sleep(15 * time.Second)
		valueWater := rand.Intn(100) + 1
		valueWind := rand.Intn(100) + 1
		postReq := domain.ElementRequest{
			Water: valueWater,
			Wind:  valueWind,
		}
		jsonValue, _ := json.Marshal(postReq)
		resp, err := http.Post("http://localhost:4000/posts", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		// Tampilkan hasil POST di terminal
		var postRes domain.ElementResponse
		json.NewDecoder(resp.Body).Decode(&postRes)
		fmt.Printf("{\n   valuewater: %d,\n   valuewind: %d\n}\nstatus water : %s\nstatus wind : %s\n\n", postRes.Water, postRes.Wind, postRes.WaterStatus, postRes.WindStatus)
	}
}