package main

import (
	"fmt"
	"os"
	"sensor-api/controller"
	"sensor-api/sensor"
	"sensor-api/weatherapi"

	"github.com/gin-gonic/gin"
)

func main() {
	weatherapi.LiveWeatherApi()
	sensor.SensorApi()

	router := gin.Default()
	router.GET("/sensor", controller.GetSensorData)
	router.GET("/weather", controller.GetWeatherData)
	router.GET("/predictions",controller.GetPredictions)
	port := os.Getenv("PORT")
	fmt.Print(port)

	if port == "" {
		port = "3000"
	}
	router.Run(":" + port)
}
