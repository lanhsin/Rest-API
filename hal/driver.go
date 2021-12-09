package hal

const (
	SENSOR_TYPE_TEMPERATURE = iota
	SENSOR_TYPE_LED
	SENSOR_TYPE_FAN
	SENSOR_TYPE_VOLTAGE
	SENSOR_TYPE_PSU
	SENSOR_TYPE_SSD
	SENSOR_TYPE_CONTROL // Reset button
	SENSOR_TYPE_END
)

const (
	SENSOR_UNIT_CEL = iota
	SENSOR_UNIT_RPM
	SENSOR_UNIT_VOLT
	SENSOR_UNIT_AMP
	SENSOR_UNIT_WATT
	SENSOR_UNIT_BOOL
	SENSOR_UNIT_PERCENT
	SENSOR_UNIT_END
)

type Extended interface {
	GetThreshold(memberId int64) (upperNonCrit float64,
		upperCirt float64,
		upperFatal float64,
		lowerNonCrit float64,
		lowerCrit float64,
		lowerFatal float64,
		err error)
	GetRange(memberId int64) (min float64, max float64, err error)
	GetPhysicalContext(memberId int64) (ctx string, err error)
	GetRelatedItem(memberId int64) (OdataId []string, err error)
}

type SensorDriver interface {
	GetNumber() (number int64, err error)
	GetName(memberId int64) (name string, err error)
	GetRaw(memberId int64) (data float64, unit string, err error)
	SetRaw(memberId int64, data float64) (err error)
	Extended
}

type FanDriver interface {
	GetNumber() (number int64, err error)
	GetName(memberId int64) (name string, err error)
	GetSpeed(memberId int64) (data int64, err error)
	SetSpeed(memberId int64, data int64) (err error)
	Extended
}

type SsdDriver interface {
	GetNumber() (number int64, err error)
	GetName(memberId int64) (name string, err error)
	GetManufacturer(memberId int64) (data string, err error)
	GetModel(memberId int64) (data string, err error)
	GetSerialNumber(memberId int64) (data string, err error)
	GetTemperature(memberId int64) (data float64, err error)
	SetReset(id int64) (err error)
}

type PsuInputRanges struct {
	InputType  string
	MinVoltage float64
	MaxVoltage float64
	MinHz      float64
	MaxHz      float64
	OutputWatt float64
}

type PsuDriver interface {
	GetNumber() (number int64, err error)
	GetName(memberId int64) (name string, err error)
	GetPowerSupplyType(memberId int64) (ty string, err error)
	GetLineInputVoltageType(memberId int64) (ty string, err error)
	GetLineInputVoltage(memberId int64) (voltage float64, err error)
	GetPowerCapacityWatts(memberId int64) (voltage float64, err error)
	GetLastPowerOutputWatts(memberId int64) (voltage float64, err error)
	GetPowerInputWatts(memberId int64) (voltage float64, err error)
	GetPowerOutputWatts(memberId int64) (voltage float64, err error)
	GetModel(memberId int64) (name string, err error)
	GetManufacturer(memberId int64) (name string, err error)
	GetFirmwareVersion(memberId int64) (name string, err error)
	GetSerialNumber(memberId int64) (name string, err error)
	GetPartNumber(memberId int64) (name string, err error)
	GetSparePartNumber(memberId int64) (name string, err error)
	GetInputRanges(memberId int64) (ranges []PsuInputRanges, err error)
}

type RpcMessager interface {
}


type PlatServer struct {
	Temperature SensorDriver
	Led         SensorDriver
	Button      SensorDriver
	Fan         FanDriver
	Ssd         SsdDriver
	Voltage     SensorDriver
	Psu         PsuDriver
}


