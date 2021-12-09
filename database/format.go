package database

const (
	// Prefix
	PrefixTemperature = "Temperature"
	PrefixLed         = "Led"
	PrefixButton      = "Button"
	PrefixSsd         = "Ssd"
	PrefixFan         = "Fan"
	PrefixVoltage     = "Voltage"
	PrefixPsu         = "Psu"

	// Extended
	ThresholdUnc    = "ThresholdUnc"
	ThresholdUc     = "ThresholdUc"
	ThresholdUf     = "ThresholdUf"
	ThresholdLnc    = "ThresholdLnc"
	ThresholdLc     = "ThresholdLc"
	ThresholdLf     = "ThresholdLf"
	RangeMin        = "RangeMin"
	RangeMax        = "RangeMax"
	PhysicalContext = "PhysicalContext"
	RelatedItem     = "RelatedItem"

	// Common
	Number       = "Number"
	MemoryId     = "MemoryId"
	Name         = "Name"
	SensorNumber = "SensorNumber"

	// Sensor
	Reading = "Reading"
	Unit    = "Unit"

	// Led
	Action = "Action"

	// Fan
	Speed       = "Speed"
	ApplySpeed  = "ApplySpeed"

	// Ssd
	Manufacturer = "Manufacturer"
	Model        = "Model"
	SerialNumber = "SerialNumber"
	Temperature  = "Temperature"

	// Voltage

	// Psu
	PowerSupplyType      = "PowerSupplyType"
	LineInputVoltageType = "LineInputVoltageType"
	LineInputVoltage     = "LineInputVoltage"
	PowerCapacityWatts   = "PowerCapacityWatts"
	LastPowerOutputWatts = "LastPowerOutputWatts"
	PowerInputWatts      = "PowerInputWatts"
	PowerOutputWatts     = "PowerOutputWatts"
	FirmwareVersion      = "FirmwareVersion"
	PartNumber           = "PartNumber"
	SparePartNumber      = "SparePartNumber"
	// Key exmaples for InputRange
	// Psu.0.InputRangeNumber
	// Psu.0.InputRange.0.InputType
	// Psu.0.InputRange.0.MinimumVoltage
	// ...
	// Psu.0.InputRange.1.InputType
	// ...
	InputRangesNumber             = "InputRangesNumber"
	InputRangesInputType          = "InputType"
	InputRangesMinimumVoltage     = "MinimumVoltage"
	InputRangesMaximumVoltage     = "MaximumVoltage"
	InputRangesMinimumFrequencyHz = "MinimumFrequencyHz"
	InputRangesMaximumFrequencyHz = "MaximumFrequencyHz"
	InputRangesOutputWattage      = "OutputWattage"
)

type ExtendedFormatter interface {
	GetThresholdUncKey(id int64) (key string)
	GetThresholdUcKey(id int64) (key string)
	GetThresholdUfKey(id int64) (key string)
	GetThresholdLncKey(id int64) (key string)
	GetThresholdLcKey(id int64) (key string)
	GetThresholdLfKey(id int64) (key string)
	GetRangeMinKey(id int64) (key string)
	GetRangeMaxKey(id int64) (key string)
	GetPhysicalContextKey(id int64) (key string)
	GetRelatedItemKey(id int64) (key string)
}

type SensorFormatter interface {
	GetNumberKey() (key string)
	GetMemoryIdKey(id int64) (key string)
	GetNameKey(id int64) (key string)
	GetSensorNumberKey(id int64) (key string)
	GetReadingKey(id int64) (key string)
	GetUnitKey(id int64) (key string)
	ExtendedFormatter
}

type LedFormatter interface {
	SensorFormatter
	GetActionKey(id int64) (key string)
}

type FanFormatter interface {
	GetNumberKey() (key string)
	GetMemoryIdKey(id int64) (key string)
	GetNameKey(id int64) (key string)
	GetSensorNumberKey(id int64) (key string)
	GetSpeedKey(id int64) (key string)
	GetApplySpeedKey(id int64) (key string)
	ExtendedFormatter
}

type SsdFormatter interface {
	GetNumberKey() (key string)
	GetMemoryIdKey(id int64) (key string)
	GetNameKey(id int64) (key string)
	GetSensorNumberKey(id int64) (key string)
	GetManufacturerKey(id int64) (key string)
	GetModelKey(id int64) (key string)
	GetSerialNumberKey(id int64) (key string)
	GetTemperatureKey(id int64) (key string)
	GetActionKey(id int64) (key string)
}

type VoltageFormatter interface {
	SensorFormatter
}

type PsuFormatter interface {
	GetNumberKey() (key string)
	GetMemoryIdKey(id int64) (key string)
	GetNameKey(id int64) (key string)
	GetPowerSupplyTypeKey(id int64) (key string)
	GetLineInputVoltageTypeKey(id int64) (key string)
	GetLineInputVoltageKey(id int64) (key string)
	GetPowerCapacityWattsKey(id int64) (key string)
	GetLastPowerOutputWattsKey(id int64) (key string)
	GetPowerInputWattsKey(id int64) (key string)
	GetPowerOutputWattsKey(id int64) (key string)
	GetModelKey(id int64) (key string)
	GetManufacturerKey(id int64) (key string)
	GetFirmwareVersionKey(id int64) (key string)
	GetSerialNumberKey(id int64) (key string)
	GetPartNumberKey(id int64) (key string)
	GetSparePartNumberKey(id int64) (key string)
	GetInputRangesNumberKey(id int64) (key string)
	GetInputRangesInputTypeKey(id int64, rangeId int64) (key string)
	GetInputRangesMinimumVoltageKey(id int64, rangeId int64) (key string)
	GetInputRangesMaximumVoltageKey(id int64, rangeId int64) (key string)
	GetInputRangesMinimumFrequencyHzKey(id int64, rangeId int64) (key string)
	GetInputRangesMaximumFrequencyHzKey(id int64, rangeId int64) (key string)
	GetInputRangesOutputWattageKey(id int64, rangeId int64) (key string)
}
