package weatherapi

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func GetHistoricalWeatherData() (*WeatherData, error) {
	baseURL := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/48858/today?unitGroup=metric&key=H39TGR9W97XMEM6NV97YSTCPW&contentType=json"

	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weatherData WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, err
	}

	return &weatherData, nil
}

func saveDataToCSV(data *WeatherData, cityName string, writer *csv.Writer) error {
	for _, day := range data.Days {
		for _, hour := range day.Hours {
			record := []string{
				cityName,
				fmt.Sprintf("%f", data.Latitude),
				fmt.Sprintf("%f", data.Longitude),
				data.Timezone,
				fmt.Sprintf("%f", data.Tzoffset),
				day.Datetime,
				hour.Datetime,
				fmt.Sprintf("%f", hour.Temp),
				fmt.Sprintf("%f", hour.Humidity),
				fmt.Sprintf("%f", hour.Precip),
				fmt.Sprintf("%f", hour.Snow),
				fmt.Sprintf("%f", hour.Pressure),
			}
			if err := writer.Write(record); err != nil {
				return err
			}
		}
	}
	return nil
}

func LiveWeatherApi() {
	cityName := "Mount Pleasant"

	data, err := GetHistoricalWeatherData()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("weather_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Writing CSV Header
	headers := []string{
		"City", "Latitude", "Longitude", "Timezone", "Timezone Offset","Date",
		"Time", "Temperature (°C)", "Humidity (%)", "Precipitation (mm)",
		"Snow (mm)", "Pressure (hPa)",
	}
	if err := writer.Write(headers); err != nil {
		log.Fatal(err)
	}

	// Writing the weather data to the CSV file
	if err := saveDataToCSV(data, cityName, writer); err != nil {
		log.Fatal(err)
	}
}