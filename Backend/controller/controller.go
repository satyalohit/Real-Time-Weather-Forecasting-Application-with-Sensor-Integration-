package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"sensor-api/sensor"
	"strings"

	"sensor-api/weatherapi"

	"github.com/gin-gonic/gin"
)

type Predictions struct {
	DateTime             string `json:"datetime"`
	PredictedTemperature string `json:"predicted_temp"`
	PredictedHumidity    string `json:"predicted_humidity"`
}

func GetWeatherData(c *gin.Context) {
	// var startDate, endDate string
	// startDate = c.Query("startDate")
	// endDate = c.Query("endDate")
	// baseURL := fmt.Sprintf("https://archive-api.open-meteo.com/v1/archive?latitude=43.5978&longitude=-84.7675&start_date=%s&end_date=%s&hourly=temperature_2m,relativehumidity_2m,dewpoint_2m,apparent_temperature,precipitation,rain,snowfall", startDate, endDate)

	// resp, err := http.Get(baseURL)
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer resp.Body.Close()

	// var weatherData weatherapi.WeatherData
	// if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
	// 	log.Println(err)
	// }
	// c.JSON(http.StatusOK, gin.H{
	// 	"data": weatherData,
	// })
	baseURL := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/48858/today?unitGroup=metric&key=H39TGR9W97XMEM6NV97YSTCPW&contentType=json"

	resp, err := http.Get(baseURL)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	var weatherData weatherapi.WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": weatherData,
	})

}

func GetSensorData(c *gin.Context) {

	res := sensor.SensorApi()

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})

}

func GetPredictions(c *gin.Context) {
	cmd := exec.Command("python3", "./data_analysis.py")
	var out, errLog bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errLog
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running script:", err)
		fmt.Println("Standard Error Output:", errLog.String())
		return
	}

	lines := strings.Split(out.String(), "\n")
	var sensorData []Predictions

	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 5 {
			fmt.Println("Unexpected line format:", line)
			continue
		}

		temp := fields[3]
		humidity := fields[4]

		sensorData = append(sensorData, Predictions{
			DateTime:             fields[1] + " " + fields[2],
			PredictedTemperature: temp,
			PredictedHumidity:    humidity,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"predictions": sensorData,
	})
}
