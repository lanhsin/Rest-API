package schema

var (
	AirFlow  = "AirFlow"
	Altitude = "Altitude"
)

type ReadingUnitsFormat struct {
	AirFlow      string
	Altitude     string
	Barometric   string
	Current      string
	EnergyJoules string
	EnergykWh    string
	Frequency    string
	Humidity     string
	LiquidFlow   string
	LiquidLevel  string
	Percent      string
	Power        string
	Pressure     string
	Rotational   string
	Temperature  string
	Voltage      string
}

var ReadingUnits ReadingUnitsFormat

func InitProperty() {
	ReadingUnits = ReadingUnitsFormat{
		AirFlow:      "cft_i",
		Altitude:     "m",
		Barometric:   "mm[Hg]",
		Current:      "A",
		EnergyJoules: "J",
		EnergykWh:    "kW.h",
		Frequency:    "Hz",
		Humidity:     "%",
		LiquidFlow:   "L/s",
		LiquidLevel:  "cm",
		Percent:      "%",
		Power:        "W",
		Pressure:     "Pa",
		Rotational:   "RPM",
		Temperature:  "Cel",
		Voltage:      "V",
	}
}
