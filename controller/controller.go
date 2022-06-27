package controller

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/Diva2504/Assignment-3-GLNG-KS03-003/models"
	"github.com/gin-gonic/gin"
)

var (
	statusWater string
	statusWind  string
	status      models.Status

// 	// mstatusWater = make(chan string)
// 	// mstatusWind  = make(chan string)
// 	// mstatus      = make(chan models.Status)
)

func RandomNumber() int {
	random := rand.Intn(100-1) + 1
	return random
}

func UpdateStatus() {
	for {
		water := RandomNumber()
		wind := RandomNumber()
		status.Water = water
		status.Wind = wind

		jsonStatus, err := json.Marshal(status)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile("status.json", jsonStatus, 0644)

		if err != nil {
			panic(err)
		}

		time.Sleep(15 * time.Second)
	}
}

func SetStatus() (models.Status, string, string) {
	fileStatus, err := ioutil.ReadFile("status.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileStatus, &status)
	if err != nil {
		panic(err)
	}

	if status.Water < 5 {
		statusWater = "Aman"
	} else if status.Water >= 6 && status.Water <= 8 {
		statusWater = "Siaga"
	} else if status.Water > 8 {
		statusWater = "Berbahaya"
	}

	if status.Wind < 6 {
		statusWind = "Aman"
	} else if status.Wind >= 7 && status.Wind <= 15 {
		statusWind = "Siaga"
	} else if status.Wind > 15 {
		statusWind = "Berbahaya"
	}
	return status, statusWater, statusWind
}

func GetStatus(c *gin.Context) {

	status, statusWater, statusWind = SetStatus()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"water":       status.Water,
		"statuswater": statusWater,
		"wind":        status.Wind,
		"statuswind":  statusWind,
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"water":       status.Water,
	// 	"statuswater": statusWater,
	// 	"wind":        status.Wind,
	// 	"statuswind":  statusWind,
	// })
}
