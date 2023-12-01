package sensor

// // type Data struct {
// // 	DeduplicationID string      `bson:"deduplicationId"`
// // 	Time            string      `bson:"time"`
// // 	DeviceInfo      DeviceInfo  `bson:"deviceInfo"`
// // 	DevAddr         string      `bson:"devAddr"`
// // 	ADR             bool        `bson:"adr"`
// // 	DR              int         `bson:"dr"`
// // 	FCnt            int         `json:"fCnt"`
// // 	FPort           int         `json:"fPort"`
// // 	Confirmed       bool        `json:"confirmed"`
// // 	RawData         string      `json:"data"`
// // 	Object          Object      `json:"object"`
// // 	RxInfo          []RxInfo    `json:"rxInfo"`
// // 	TxInfo          TxInfo      `json:"txInfo"`
// // }

// // type DeviceInfo struct {
// // 	TenantID          string            `json:"tenantId"`
// // 	TenantName        string            `json:"tenantName"`
// // 	ApplicationID     string            `json:"applicationId"`
// // 	ApplicationName   string            `json:"applicationName"`
// // 	DeviceProfileID   string            `json:"deviceProfileId"`
// // 	DeviceProfileName string            `json:"deviceProfileName"`
// // 	DeviceName        string            `json:"deviceName"`
// // 	DevEui            string            `json:"devEui"`
// // 	Tags              map[string]string `json:"tags"`
// // }

// // type Object struct {
// // 	DeviceID               float64            `json:"device_id"`
// // 	BatteryVoltage         Measurement        `json:"battery_voltage"`
// // 	BarometerTemperature   Measurement        `json:"barometer_temperature"`
// // 	RawIRReadingLPF        Measurement        `json:"raw_ir_reading_lpf"`
// // 	AirTemperature         Measurement        `json:"air_temperature"`
// // 	ProtocolVersion        float64            `json:"protocol_version"`
// // 	AirHumidity            Measurement        `json:"air_humidity"`
// // 	CO2SensorStatus        Measurement        `json:"co2_sensor_status"`
// // 	CO2Concentration       Measurement        `json:"co2_concentration"`
// // 	BarometricPressure     Measurement        `json:"barometric_pressure"`
// // 	RawIRReading           Measurement        `json:"raw_ir_reading"`
// // 	CapacitorVoltage2      Measurement        `json:"capacitor_voltage_2"`
// // 	CO2ConcentrationLPF    Measurement        `json:"co2_concentration_lpf"`
// // 	CapacitorVoltage1      Measurement        `json:"capacitor_voltage_1"`
// // 	CO2SensorTemperature   Measurement        `json:"co2_sensor_temperature"`
// // }

// // type Measurement struct {
// // 	Unit       string  `json:"unit"`
// // 	DisplayName string `json:"displayName"`
// // 	Value      float64 `json:"value"`
// // }

// // type RxInfo struct {
// // 	GatewayID        string     `json:"gatewayId"`
// // 	UplinkID         int        `json:"uplinkId"`
// // 	Time             string     `json:"time"`
// // 	TimeSinceGPSEpoch string     `json:"timeSinceGpsEpoch"`
// // 	Rssi             int        `json:"rssi"`
// // 	Snr              float64    `json:"snr"`
// // 	Location         Location   `json:"location"`
// // 	Context          string     `json:"context"`
// // 	Metadata         Metadata   `json:"metadata"`
// // 	CRCStatus        string     `json:"crcStatus"`
// // }

// // type Location struct {
// // 	Latitude  float64 `json:"latitude"`
// // 	Longitude float64 `json:"longitude"`
// // 	Altitude  float64 `json:"altitude"`
// // }

// // type Metadata struct {
// // 	RegionCommonName string `json:"region_common_name"`
// // 	RegionConfigID   string `json:"region_config_id"`
// // }

// // type TxInfo struct {
// // 	Frequency  int        `json:"frequency"`
// // 	Modulation Modulation `json:"modulation"`
// // }

// // type Modulation struct {
// // 	Lora Lora `json:"lora"`
// // }

// // type Lora struct {
// // 	Bandwidth      int    `json:"bandwidth"`
// // 	SpreadingFactor int   `json:"spreadingFactor"`
// // 	CodeRate       string `json:"codeRate"`
// // }

type Data struct {
	Date                   string
	TimeOfDay              string
	Time       string     `json:"time" bson:"time"`
	DeviceInfo DeviceInfo `json:"deviceInfo" bson:"deviceInfo"`
	Object     Object     `json:"object" bson:"object"`
}

type DeviceInfo struct {
	DeviceProfileName string `json:"deviceProfileName" bson:"deviceProfileName"`
}

type Object struct {
	AirTemperature       Measurement `json:"air_temperature" bson:"air_temperature"`
	AirHumidity          Measurement `json:"air_humidity" bson:"air_humidity"`
	BarometricPressure   Measurement `json:"barometric_pressure" bson:"barometric_pressure"`
	CO2SensorTemperature Measurement `json:"co2_sensor_temperature" bson:"co2_sensor_temperature"`
}

type Measurement struct {
	Unit        string  `json:"unit" bson:"unit"`
	DisplayName string  `json:"displayName" bson:"displayName"`
	Value       float64 `json:"value" bson:"value"`
}
type Sdata struct {
	Date                   string
	TimeOfDay              string
	Temp                   float64
	Humidity               float64
	Barometric_pressure    int
	Co2_sensor_temperature float64
}

// type Data struct {
// 	DeduplicationID string    `json:"deduplicationId" bson:"deduplicationId"`
// 	Time            time.Time `json:"time" bson:"time`
// 	DeviceInfo      struct {
// 		DeviceProfileName string `json:"deviceProfileName" bson:"deviceProfileName"`
// 	} `json:"deviceInfo" bson:"deviceInfo"`
// 	Object struct {
// 		DeviceID       float64 `json:"device_id" bson:"device_id"`
// 		AirTemperature struct {
// 			Unit        string  `json:"unit" bson:"unit"`
// 			Value       float64 `json:"value" bson:"value"`
// 			DisplayName string  `json:"displayName" bson:"displayName"`
// 		} `json:"air_temperature" bson:"air_temperature"`
// 		AirHumidity struct {
// 			Value       float64 `json:"value" bson:"value"`
// 			DisplayName string  `json:"displayName" bson:"displayName"`
// 			Unit        string  `json:"unit" bson:"unit"`
// 		} `json:"air_humidity" bson:"air_humidity"`
// 		BarometricPressure struct {
// 			Unit        string  `json:"unit" bson:"unit"`
// 			DisplayName string  `json:"unit" bson:"unit" `
// 			Value       float64 `json:"value" bson:"value"`
// 		} `json:"barometric_pressure" bson:"barometric_pressure"`

// 		Co2SensorTemperature struct {
// 			Value       float64 `json:"value" bson:"value"`
// 			Unit        string  `json:"unit" bson:"unit"`
// 			DisplayName string  `json:"displayName" bson:"displayName"`
// 		} `json:"co2_sensor_temperature" bson:"co2_sensor_temperature"`
// 	} `json:"object" bson:"object" `
// }
