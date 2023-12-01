package weatherapi



// type WeatherData struct {
// 	Latitude             float64 `json:"latitude"`
// 	Longitude            float64 `json:"longitude"`
// 	GenerationtimeMs     float64 `json:"generationtime_ms"`
// 	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
// 	Timezone             string  `json:"timezone"`
// 	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
// 	Elevation            float64 `json:"elevation"`
// 	HourlyUnits          struct {
// 		Time                string `json:"time"`
// 		Temperature2M       string `json:"temperature_2m"`
// 		Relativehumidity2M  string `json:"relativehumidity_2m"`
// 		Dewpoint2M          string `json:"dewpoint_2m"`
// 		ApparentTemperature string `json:"apparent_temperature"`
// 		Precipitation       string `json:"precipitation"`
// 		Rain                string `json:"rain"`
// 		Snowfall            string `json:"snowfall"`
// 	} `json:"hourly_units"`
// 	Hourly struct {
// 		Time                []string      `json:"time"`
// 		Temperature2M       []interface{} `json:"temperature_2m"`
// 		Relativehumidity2M  []interface{} `json:"relativehumidity_2m"`
// 		Dewpoint2M          []interface{} `json:"dewpoint_2m"`
// 		ApparentTemperature []interface{} `json:"apparent_temperature"`
// 		Precipitation       []interface{} `json:"precipitation"`
// 		Rain                []interface{} `json:"rain"`
// 		Snowfall            []interface{} `json:"snowfall"`
// 	} `json:"hourly"`
// }
type WeatherData struct {
	QueryCost       int     `json:"queryCost"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	ResolvedAddress string  `json:"resolvedAddress"`
	Address         string  `json:"address"`
	Timezone        string  `json:"timezone"`
	Tzoffset        float64 `json:"tzoffset"`
	Days            []struct {
		Datetime      string  `json:"datetime"`
		DatetimeEpoch int     `json:"datetimeEpoch"`
		Temp          float64 `json:"temp"`
		Humidity      float64 `json:"humidity"`
		Precip        float64 `json:"precip"`
		Snow          float64 `json:"snow"`
		Pressure      float64 `json:"pressure"`
		Hours         []struct {
			Datetime      string  `json:"datetime"`
			DatetimeEpoch int     `json:"datetimeEpoch"`
			Temp          float64 `json:"temp"`
			Humidity      float64 `json:"humidity"`
			Precip        float64 `json:"precip"`
			Snow          float64 `json:"snow"`
			Pressure      float64 `json:"pressure"`
		} `json:"hours"`
	} `json:"days"`
	Alerts            []any `json:"alerts"`
	CurrentConditions struct {
		Datetime      string  `json:"datetime"`
		DatetimeEpoch int     `json:"datetimeEpoch"`
		Temp          float64 `json:"temp"`
		Humidity      float64 `json:"humidity"`
		Precip        float64 `json:"precip"`
		Snow          float64 `json:"snow"`
		Pressure      float64 `json:"pressure"`
	} `json:"currentConditions"`
}