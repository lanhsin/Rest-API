package hal

import (
	"errors"
	"pemgr/cmd/pemgr-server/schema"
	"strconv"
)

type MockupTemperature struct {
}

func (mock *MockupTemperature) GetNumber() (int64, error) {
	return 7, nil
}

func (mock *MockupTemperature) GetName(id int64) (string, error) {
	name := ""
	switch id {
	case 0:
		name = "Mockup SWM Local Temperature"
	case 1:
		name = "Mockup SWM Remote Temperature"
	case 2:
		name = "Mockup DIMM Temperature"
	case 3:
		name = "Mockup ODP Room Local Temperature 1"
	case 4:
		name = "Mockup ODP Room Remote Temperature 1"
	case 5:
		name = "Mockup ODP Room Local Temperature 2"
	case 6:
		name = "Mockup ODP Room Remote Temperature 2"
	default:
		return name, errors.New("Unknown id")
	}
	return name, nil
}

func (mock *MockupTemperature) GetRaw(id int64) (float64, string, error) {
	return float64(30 + id), schema.ReadingUnits.Temperature, nil
}

func (mock *MockupTemperature) SetRaw(id int64, raw float64) error {
	return nil
}

func (mock *MockupTemperature) GetThreshold(id int64) (upperNonCrit float64,
	upperCrit float64,
	upperFatal float64,
	lowerNonCrit float64,
	lowerCrit float64,
	lowerFatal float64,
	err error) {
	return 0, 0, 0, 0, 0, 0, nil
}

func (mock *MockupTemperature) GetRange(id int64) (min float64, max float64, err error) {
	return 0, 0, nil
}

func (mock *MockupTemperature) GetPhysicalContext(id int64) (string, error) {
	ctx := ""
	switch id {
	case 0, 1:
		ctx = "SystemBoard"
	case 2:
		ctx = "Memory"
	case 3, 4, 5, 6:
		ctx = "Backplane"
	default:
		return ctx, errors.New("Unknown id")
	}
	return ctx, nil
}

func (mock *MockupTemperature) GetRelatedItem(memberId int64) (OdataId []string, err error) {
	var oids []string
	return oids, nil
}

type MockupLed struct {
	Light [1024]float64
}

func (mock *MockupLed) GetNumber() (int64, error) {
	return 27, nil
}

func (mock *MockupLed) GetName(id int64) (string, error) {
	name := ""
	switch id {
	case 0:
		name = "Mockup System Location LED"
	case 1:
		name = "Mockup System Fault LED"
	case 2:
		name = "Mockup Segments LED A1"
	case 3:
		name = "Mockup Segments LED B1"
	case 4:
		name = "Mockup Segments LED C1"
	case 5:
		name = "Mockup Segments LED D1"
	case 6:
		name = "Mockup Segments LED E1"
	case 7:
		name = "Mockup Segments LED F1"
	case 8:
		name = "Mockup Segments LED G1"
	case 9:
		name = "Mockup Segments LED DP1"
	case 10:
		name = "Mockup Segments LED A2"
	case 11:
		name = "Mockup Segments LED B2"
	case 12:
		name = "Mockup Segments LED C2"
	case 13:
		name = "Mockup Segments LED D2"
	case 14:
		name = "Mockup Segments LED E2"
	case 15:
		name = "Mockup Segments LED F2"
	case 16:
		name = "Mockup Segments LED G2"
	case 17:
		name = "Mockup Segments LED DP2"
	case 18:
		name = "IOM Attention DIMM LED"
	case 19:
		name = "IOM Attention FAN1 LED"
	case 20:
		name = "IOM Attention FAN2 LED"
	case 21:
		name = "IOM Attention FAN3 LED"
	case 22:
		name = "IOM Attention FAN4 LED"
	case 23:
		name = "IOM Attention FAN5 LED"
	case 24:
		name = "IOM Attention FAN6 LED"
	case 25:
		name = "IOM Attention NVMe LED"
	case 26:
		name = "IOM Attention VBat LED"
	default:
		return name, errors.New("Unknown id")
	}
	return name, nil
}

func (mock *MockupLed) GetRaw(id int64) (float64, string, error) {
	return mock.Light[id], "Boolean", nil
}

func (mock *MockupLed) SetRaw(id int64, raw float64) error {
	mock.Light[id] = raw
	return nil
}

func (mock *MockupLed) GetThreshold(id int64) (upperNonCrit float64,
	upperCrit float64,
	upperFatal float64,
	lowerNonCrit float64,
	lowerCrit float64,
	lowerFatal float64,
	err error) {
	return 0, 0, 0, 0, 0, 0, nil
}

func (mock *MockupLed) GetRange(id int64) (min float64, max float64, err error) {
	return 0, 0, nil
}

func (mock *MockupLed) GetPhysicalContext(id int64) (string, error) {
	ctx := ""
	switch {
	case id >= int64(0) && id <= int64(26):
		ctx = "Front"
	default:
		return ctx, errors.New("Unknown id")
	}
	return ctx, nil
}

func (mock *MockupLed) GetRelatedItem(id int64) (OdataId []string, err error) {
	var oids []string
	return oids, errors.New("No related item")
}

type MockupButton struct {
}

func (mock *MockupButton) GetNumber() (int64, error) {
	return 1, nil
}

func (mock *MockupButton) GetName(id int64) (string, error) {
	name := ""
	switch id {
	case 0:
		name = "Mockup ODP Reset Button"
	default:
		return name, errors.New("Unknown id")
	}
	return name, nil
}

func (mock *MockupButton) GetRaw(id int64) (float64, string, error) {
	return float64(id % 2), "Boolean", nil
}

func (mock *MockupButton) SetRaw(id int64, raw float64) error {
	return nil
}

func (mock *MockupButton) GetThreshold(id int64) (upperNonCrit float64,
	upperCrit float64,
	upperFatal float64,
	lowerNonCrit float64,
	lowerCrit float64,
	lowerFatal float64,
	err error) {
	return 0, 0, 0, 0, 0, 0, nil
}

func (mock *MockupButton) GetRange(memberId int64) (min float64, max float64, err error) {
	return 0, 0, nil
}

func (mock *MockupButton) GetPhysicalContext(id int64) (string, error) {
	ctx := ""
	switch id {
	case 0:
		ctx = "Front"
	default:
		return ctx, errors.New("Unknown id")
	}
	return ctx, nil
}

func (mock *MockupButton) GetRelatedItem(memberId int64) (OdataId []string, err error) {
	var oids []string
	return oids, errors.New("No related item")
}

type MockupFan struct {
	Speed [64]int64
}

func (mock *MockupFan) GetNumber() (int64, error) {
	return 12, nil
}

func (mock *MockupFan) GetName(id int64) (string, error) {
	name := ""
	switch id {
	case 0:
		name = "Mokup INLET FAN 1"
	case 1:
		name = "Mokup INLET FAN 2"
	case 2:
		name = "Mokup INLET FAN 3"
	case 3:
		name = "Mokup INLET FAN 4"
	case 4:
		name = "Mokup INLET FAN 5"
	case 5:
		name = "Mokup INLET FAN 6"
	case 6:
		name = "Mokup OUTLET FAN 1"
	case 7:
		name = "Mokup OUTLET FAN 2"
	case 8:
		name = "Mokup OUTLET FAN 3"
	case 9:
		name = "Mokup OUTLET FAN 4"
	case 10:
		name = "Mokup OUTLET FAN 5"
	case 11:
		name = "Mokup OUTLET FAN 6"
	default:
		return name, errors.New("Unknown id")
	}
	return name, nil
}

func (mock *MockupFan) GetSpeed(id int64) (int64, error) {
	return int64(mock.Speed[id]), nil
}

func (mock *MockupFan) SetSpeed(id int64, speed int64) error {
	mock.Speed[id] = speed
	return nil
}

func (mock *MockupFan) GetThreshold(id int64) (upperNonCrit float64,
	upperCrit float64,
	upperFatal float64,
	lowerNonCrit float64,
	lowerCrit float64,
	lowerFatal float64,
	err error) {
	return 0, 0, 0, 0, 0, 0, nil
}

func (mock *MockupFan) GetRange(memberId int64) (min float64, max float64, err error) {
	return 0, 0, nil
}

func (mock *MockupFan) GetPhysicalContext(id int64) (string, error) {
	ctx := ""
	switch {
	case id >= 0 && id <= 11:
		ctx = "Fan"
	default:
		return ctx, errors.New("Unknown id")
	}
	return ctx, nil
}

func (mock *MockupFan) GetRelatedItem(memberId int64) (OdataId []string, err error) {
	var oids []string
	return oids, errors.New("No related item")
}

type MockupSsd struct {
}

func (mock *MockupSsd) GetNumber() (int64, error) {
	return 24, nil
}

func (mock *MockupSsd) GetName(id int64) (string, error) {
	return "Mockup SSD " + strconv.FormatInt(id+1, 10), nil
}

func (mock *MockupSsd) GetManufacturer(id int64) (string, error) {
	idStr := strconv.FormatInt(id, 10)
	return "Mockup Manufacturer " + idStr, nil
}

func (mock *MockupSsd) GetModel(id int64) (string, error) {
	idStr := strconv.FormatInt(id, 10)
	return "Mockup Model " + idStr, nil
}

func (mock *MockupSsd) GetSerialNumber(id int64) (string, error) {
	idStr := strconv.FormatInt(id, 10)
	return "Mokup SerialNumber " + idStr, nil
}

func (mock *MockupSsd) GetTemperature(id int64) (float64, error) {
	return float64(30 + id), nil
}

func (mock *MockupSsd) SetReset(id int64) error {
	return nil
}

type MockupVoltage struct {
}

func (mock *MockupVoltage) GetNumber() (int64, error) {
	return 5, nil
}

func (mock *MockupVoltage) GetName(id int64) (string, error) {
	name := ""
	switch id {
	case 0:
		name = "Mockup Voltage 1"
	case 1:
		name = "Mockup Voltage 2"
	case 2:
		name = "Mockup Voltage 3"
	case 3:
		name = "Mockup Voltage 4"
	case 4:
		name = "Mockup Voltage 5"
	case 5:
		name = "Mockup Voltage 6"
	default:
		return name, errors.New("Unknown id")
	}
	return name, nil
}

func (mock *MockupVoltage) GetRaw(id int64) (float64, string, error) {
	return float64(5 + id), schema.ReadingUnits.Voltage, nil
}

func (mock *MockupVoltage) SetRaw(id int64, raw float64) error {
	return nil
}

func (mock *MockupVoltage) GetThreshold(id int64) (upperNonCrit float64,
	upperCrit float64,
	upperFatal float64,
	lowerNonCrit float64,
	lowerCrit float64,
	lowerFatal float64,
	err error) {
	return 0, 0, 0, 0, 0, 0, nil
}

func (mock *MockupVoltage) GetRange(id int64) (min float64, max float64, err error) {
	return 0, 0, nil
}

func (mock *MockupVoltage) GetPhysicalContext(id int64) (string, error) {
	ctx := ""
	switch id {
	case 0, 1, 2, 3, 4, 5:
		ctx = "VoltageRegulator"
	default:
		return ctx, errors.New("Unknown id")
	}
	return ctx, nil
}

func (mock *MockupVoltage) GetRelatedItem(memberId int64) (OdataId []string, err error) {
	var oids []string
	return oids, nil
}

type MockupPsu struct {
}

func (mock *MockupPsu) GetNumber() (int64, error) {
	return 2, nil
}

func (mock *MockupPsu) GetName(id int64) (string, error) {
	return "Mockup Psu" + strconv.FormatInt(id+1, 10), nil
}

func (mock *MockupPsu) GetPowerSupplyType(id int64) (ty string, err error) {
	return "AC", nil
}

func (mock *MockupPsu) GetLineInputVoltageType(id int64) (ty string, err error) {
	return "ACWideRange", nil
}

func (mock *MockupPsu) GetLineInputVoltage(id int64) (voltage float64, err error) {
	return 8.5 + float64(id), nil
}

func (mock *MockupPsu) GetPowerCapacityWatts(id int64) (voltage float64, err error) {
	return 9.5 + float64(id), nil
}

func (mock *MockupPsu) GetLastPowerOutputWatts(id int64) (voltage float64, err error) {
	return 10.5 + float64(id), nil
}

func (mock *MockupPsu) GetPowerInputWatts(id int64) (voltage float64, err error) {
	return 11.5 + float64(id), nil
}

func (mock *MockupPsu) GetPowerOutputWatts(id int64) (voltage float64, err error) {
	return 12.5 + float64(id), nil
}

func (mock *MockupPsu) GetModel(id int64) (name string, err error) {
	return "Mockup Model", nil
}

func (mock *MockupPsu) GetManufacturer(id int64) (name string, err error) {
	return "Mockup Manufacturer", nil
}

func (mock *MockupPsu) GetFirmwareVersion(id int64) (name string, err error) {
	return "Mockup FirmwareVersion", nil
}

func (mock *MockupPsu) GetSerialNumber(id int64) (name string, err error) {
	return "Mockup SerialNumber", nil
}

func (mock *MockupPsu) GetPartNumber(id int64) (name string, err error) {
	return "Mockup PartNumber", nil
}

func (mock *MockupPsu) GetSparePartNumber(id int64) (name string, err error) {
	return "Mockup SpareSerialNumber", nil
}

func (mock *MockupPsu) GetInputRanges(id int64) (ranges []PsuInputRanges, err error) {
	r := []PsuInputRanges{}
	for i := 0; i < 2; i++ {
		n := PsuInputRanges{InputType: "AC", MinVoltage: 1.0 + float64(i), MaxVoltage: 15.0 + float64(i),
			MinHz: 50, MaxHz: 60, OutputWatt: 100}
		r = append(r, n)
	}
	return r, nil
}
