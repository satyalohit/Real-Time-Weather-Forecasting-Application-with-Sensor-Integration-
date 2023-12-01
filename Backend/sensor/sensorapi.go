package sensor

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"os"
	"time"

	// "encoding/csv"
	"fmt"
	"log"
	"sort"

	// "os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type Response struct {
	Id          string `bson:"_id"`
	LoraContent string `bson:"loraContent"`
}

type Document struct {
	Doc []Response `bson:"documents"`
}

func SensorApi() []Sdata {

	uri := "mongodb+srv://cnets-user1:CMUcmich2023@cnets-0.jjpntwr.mongodb.net/?retryWrites=true&w=majority&appName=AtlasApp"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.TODO())

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected and pinged.")

	collection := client.Database("decentlab_sensors").Collection("readings")

	// filter := bson.M{"time",}
	

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {

		log.Fatal(err)
	}
	var documents []Response
	for cur.Next(context.TODO()) {
		var elem Response
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		

		documents = append(documents, elem)
	}

	if err := cur.Err(); err != nil {

		log.Fatal(err)
	}
	cur.Close(context.TODO())


	sort.Slice(documents, func(i, j int) bool {
        timeI, errI := parseTime(documents[i].LoraContent)
        timeJ, errJ := parseTime(documents[j].LoraContent)
        if errI != nil || errJ != nil {
            return false
        }
        return timeI.After(timeJ) // Sorting in descending order
    })

	layout := time.RFC3339
	location, _ := time.LoadLocation("America/New_York")

	var lData []Data
	var sData []Sdata
	for _, doc := range documents {
		var sdata Sdata
		var ldata Data
		err = json.Unmarshal([]byte(doc.LoraContent), &ldata)
		if err != nil {
			fmt.Println(err)
			continue
		}
		gmtTime, err := time.Parse(layout, ldata.Time)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			continue
		}
		estTime := gmtTime.In(location)
		ldata.Date = estTime.Format("2006-01-02")
		ldata.TimeOfDay = estTime.Format("15:04:05")
		sdata.Barometric_pressure = int(ldata.Object.BarometricPressure.Value)
		sdata.Co2_sensor_temperature = ldata.Object.CO2SensorTemperature.Value
		sdata.Date = ldata.Date
		sdata.TimeOfDay = ldata.TimeOfDay
		sdata.Humidity = ldata.Object.AirHumidity.Value
		sdata.Temp = ldata.Object.AirTemperature.Value
		sData = append(sData, sdata)
		lData = append(lData, ldata)
		fmt.Println(sdata.Date)

	}

	if err := ConvertToCsv(lData); err != nil {
		log.Fatal(err)
	}
	return sData

}
func parseTime(loraContent string) (time.Time, error) {
    var data struct {
        Time string `json:"time"`
    }
    err := json.Unmarshal([]byte(loraContent), &data)
    if err != nil {
        return time.Time{}, err
    }
    return time.Parse(time.RFC3339, data.Time)
}

func ConvertToCsv(data []Data) error {
	file, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{
		"DeviceProfileName",
		"Date",
		"Time",
		"AirTemperature",
		"AirHumidity",
		"BarometricPressure",
		"CO2SensorTemperature",
		"Unit",
	}
	if err := writer.Write(header); err != nil {
		fmt.Println("Error writing header to CSV:", err)
		return err
	}

	for _, dataItem := range data {
		record := []string{
			dataItem.DeviceInfo.DeviceProfileName,
			dataItem.Date,
			dataItem.TimeOfDay,
			fmt.Sprintf("%v", dataItem.Object.AirTemperature.Value),
			fmt.Sprintf("%v", dataItem.Object.AirHumidity.Value),
			fmt.Sprintf("%v", dataItem.Object.BarometricPressure.Value),
			fmt.Sprintf("%v", dataItem.Object.CO2SensorTemperature.Value),
			dataItem.Object.AirTemperature.Unit,
		}
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing record to CSV:", err)
			return err
		}
	}

	writer.Flush()
	return writer.Error()
}
