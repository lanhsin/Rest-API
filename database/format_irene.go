package database

import (
	"fmt"
)

type TemperatureFormat struct {
}

func (db *TemperatureFormat) GetNumberKey() string {
	return fmt.Sprintf("%s.%s", PrefixTemperature, Number)
}

func (db *TemperatureFormat) GetMemoryIdKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixTemperature, id, MemoryId)
}

func (db *TemperatureFormat) GetNameKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixTemperature, id, Name)
}

func (db *TemperatureFormat) GetSensorNumberKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixTemperature, id, SensorNumber)
}

func (db *TemperatureFormat) GetReadingKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixTemperature, id, Reading)
}

func (db *TemperatureFormat) GetUnitKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixTemperature, id, Unit)
}

func (db *TemperatureFormat) GetThresholdUncKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixTemperature, id, ThresholdUnc)
}

func (db *TemperatureFormat) GetThresholdUcKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixTemperature, id, ThresholdUc)
}

func (db *TemperatureFormat) GetThresholdUfKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixTemperature, id, ThresholdUf)
}

func (db *TemperatureFormat) GetThresholdLncKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixTemperature, id, ThresholdLnc)
}

func (db *TemperatureFormat) GetThresholdLcKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixTemperature, id, ThresholdLc)
}

func (db *TemperatureFormat) GetThresholdLfKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixTemperature, id, ThresholdLf)
}

func (db *TemperatureFormat) GetRangeMinKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixTemperature, id, RangeMin)
}

func (db *TemperatureFormat) GetRangeMaxKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixTemperature, id, RangeMax)
}

func (db *TemperatureFormat) GetPhysicalContextKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixTemperature, id, PhysicalContext)
}

func (db *TemperatureFormat) GetRelatedItemKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixTemperature, id, RelatedItem)
}

type LedFormat struct {
}

func (db *LedFormat) GetNumberKey() string {
	return fmt.Sprintf("%s.%s", PrefixLed, Number)
}

func (db *LedFormat) GetMemoryIdKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixLed, id, MemoryId)
}

func (db *LedFormat) GetNameKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixLed, id, Name)
}

func (db *LedFormat) GetSensorNumberKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixLed, id, SensorNumber)
}

func (db *LedFormat) GetReadingKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixLed, id, Reading)
}

func (db *LedFormat) GetUnitKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixLed, id, Unit)
}

func (db *LedFormat) GetActionKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixLed, id, Action)
}

func (db *LedFormat) GetThresholdUncKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixLed, id, ThresholdUnc)
}

func (db *LedFormat) GetThresholdUcKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixLed, id, ThresholdUc)
}

func (db *LedFormat) GetThresholdUfKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixLed, id, ThresholdUf)
}

func (db *LedFormat) GetThresholdLncKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixLed, id, ThresholdLnc)
}

func (db *LedFormat) GetThresholdLcKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixLed, id, ThresholdLc)
}

func (db *LedFormat) GetThresholdLfKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixLed, id, ThresholdLf)
}

func (db *LedFormat) GetRangeMinKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixLed, id, RangeMin)
}

func (db *LedFormat) GetRangeMaxKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixLed, id, RangeMax)
}

func (db *LedFormat) GetPhysicalContextKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixLed, id, PhysicalContext)
}

func (db *LedFormat) GetRelatedItemKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixLed, id, RelatedItem)
}

type ButtonFormat struct {
}

func (db *ButtonFormat) GetNumberKey() string {
	return fmt.Sprintf("%s.%s", PrefixButton, Number)
}

func (db *ButtonFormat) GetMemoryIdKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixButton, id, MemoryId)
}

func (db *ButtonFormat) GetNameKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixButton, id, Name)
}

func (db *ButtonFormat) GetSensorNumberKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixButton, id, SensorNumber)
}

func (db *ButtonFormat) GetReadingKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixButton, id, Reading)
}

func (db *ButtonFormat) GetUnitKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixButton, id, Unit)
}

func (db *ButtonFormat) GetThresholdUncKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixButton, id, ThresholdUnc)
}

func (db *ButtonFormat) GetThresholdUcKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixButton, id, ThresholdUc)
}

func (db *ButtonFormat) GetThresholdUfKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixButton, id, ThresholdUf)
}

func (db *ButtonFormat) GetThresholdLncKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixButton, id, ThresholdLnc)
}

func (db *ButtonFormat) GetThresholdLcKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixButton, id, ThresholdLc)
}

func (db *ButtonFormat) GetThresholdLfKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixButton, id, ThresholdLf)
}

func (db *ButtonFormat) GetRangeMinKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixButton, id, RangeMin)
}

func (db *ButtonFormat) GetRangeMaxKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixButton, id, RangeMax)
}

func (db *ButtonFormat) GetPhysicalContextKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixButton, id, PhysicalContext)
}

func (db *ButtonFormat) GetRelatedItemKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixButton, id, RelatedItem)
}

type FanFormat struct {
}

func (db *FanFormat) GetNumberKey() string {
	return fmt.Sprintf("%s.%s", PrefixFan, Number)
}

func (db *FanFormat) GetMemoryIdKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixFan, id, MemoryId)
}

func (db *FanFormat) GetNameKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixFan, id, Name)
}

func (db *FanFormat) GetSensorNumberKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixFan, id, SensorNumber)
}

func (db *FanFormat) GetSpeedKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixFan, id, Speed)
}

func (db *FanFormat) GetApplySpeedKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixFan, id, ApplySpeed)
}

func (db *FanFormat) GetThresholdUncKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixFan, id, ThresholdUnc)
}

func (db *FanFormat) GetThresholdUcKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixFan, id, ThresholdUc)
}

func (db *FanFormat) GetThresholdUfKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixFan, id, ThresholdUf)
}

func (db *FanFormat) GetThresholdLncKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixFan, id, ThresholdLnc)
}

func (db *FanFormat) GetThresholdLcKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixFan, id, ThresholdLc)
}

func (db *FanFormat) GetThresholdLfKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixFan, id, ThresholdLf)
}

func (db *FanFormat) GetRangeMinKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixFan, id, RangeMin)
}

func (db *FanFormat) GetRangeMaxKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixFan, id, RangeMax)
}

func (db *FanFormat) GetPhysicalContextKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixFan, id, PhysicalContext)
}

func (db *FanFormat) GetRelatedItemKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixFan, id, RelatedItem)
}

type SsdFormat struct {
}

func (db *SsdFormat) GetNumberKey() string {
	return fmt.Sprintf("%s.%s", PrefixSsd, Number)
}

func (db *SsdFormat) GetMemoryIdKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixSsd, id, MemoryId)
}

func (db *SsdFormat) GetNameKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixSsd, id, Name)
}

func (db *SsdFormat) GetSensorNumberKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixSsd, id, SensorNumber)
}

func (db *SsdFormat) GetManufacturerKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixSsd, id, Manufacturer)
}

func (db *SsdFormat) GetModelKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixSsd, id, Model)
}

func (db *SsdFormat) GetSerialNumberKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixSsd, id, SerialNumber)
}

func (db *SsdFormat) GetTemperatureKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixSsd, id, Temperature)
}

func (db *SsdFormat) GetActionKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixSsd, id, Action)
}

type VoltageFormat struct {
}

func (db *VoltageFormat) GetNumberKey() string {
	return fmt.Sprintf("%s.%s", PrefixVoltage, Number)
}

func (db *VoltageFormat) GetMemoryIdKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixVoltage, id, MemoryId)
}

func (db *VoltageFormat) GetNameKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixVoltage, id, Name)
}

func (db *VoltageFormat) GetSensorNumberKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixVoltage, id, SensorNumber)
}

func (db *VoltageFormat) GetReadingKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixVoltage, id, Reading)
}

func (db *VoltageFormat) GetUnitKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixVoltage, id, Unit)
}

func (db *VoltageFormat) GetThresholdUncKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixVoltage, id, ThresholdUnc)
}

func (db *VoltageFormat) GetThresholdUcKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixVoltage, id, ThresholdUc)
}

func (db *VoltageFormat) GetThresholdUfKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixVoltage, id, ThresholdUf)
}

func (db *VoltageFormat) GetThresholdLncKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixVoltage, id, ThresholdLnc)
}

func (db *VoltageFormat) GetThresholdLcKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixVoltage, id, ThresholdLc)
}

func (db *VoltageFormat) GetThresholdLfKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixVoltage, id, ThresholdLf)
}

func (db *VoltageFormat) GetRangeMinKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixVoltage, id, RangeMin)
}

func (db *VoltageFormat) GetRangeMaxKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixVoltage, id, RangeMax)
}

func (db *VoltageFormat) GetPhysicalContextKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixVoltage, id, PhysicalContext)
}

func (db *VoltageFormat) GetRelatedItemKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixVoltage, id, RelatedItem)
}

type PsuFormat struct {
}

func (db *PsuFormat) GetNumberKey() string {
	return fmt.Sprintf("%s.%s", PrefixPsu, Number)
}

func (db *PsuFormat) GetMemoryIdKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixPsu, id, MemoryId)
}

func (db *PsuFormat) GetNameKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixPsu, id, Name)
}

func (db *PsuFormat) GetPowerSupplyTypeKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixPsu, id, PowerSupplyType)
}

func (db *PsuFormat) GetLineInputVoltageTypeKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixPsu, id, LineInputVoltageType)
}

func (db *PsuFormat) GetLineInputVoltageKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixPsu, id, LineInputVoltage)
}

func (db *PsuFormat) GetPowerCapacityWattsKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixPsu, id, PowerCapacityWatts)
}

func (db *PsuFormat) GetLastPowerOutputWattsKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixPsu, id, LastPowerOutputWatts)
}

func (db *PsuFormat) GetPowerInputWattsKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixPsu, id, PowerInputWatts)
}

func (db *PsuFormat) GetPowerOutputWattsKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixPsu, id, PowerOutputWatts)
}

func (db *PsuFormat) GetModelKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixPsu, id, Model)
}

func (db *PsuFormat) GetManufacturerKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixPsu, id, Manufacturer)
}

func (db *PsuFormat) GetFirmwareVersionKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixPsu, id, FirmwareVersion)
}

func (db *PsuFormat) GetSerialNumberKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixPsu, id, SerialNumber)
}

func (db *PsuFormat) GetPartNumberKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixPsu, id, PartNumber)
}

func (db *PsuFormat) GetSparePartNumberKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixPsu, id, SparePartNumber)
}

func (db *PsuFormat) GetInputRangesNumberKey(id int64) string {
	return fmt.Sprintf("%s.%d.%s", PrefixPsu, id, InputRangesNumber)
}

func (db *PsuFormat) GetInputRangesInputTypeKey(id int64, rangeId int64) string {
	// Ex. Psu.0.InputRange.0.InputType
	return fmt.Sprintf("%s.%d.InputRange.%d.%s", PrefixPsu, id, rangeId, InputRangesInputType)
}

func (db *PsuFormat) GetInputRangesMinimumVoltageKey(id int64, rangeId int64) string {
	// Ex. Psu.0.InputRang.0.MinimumVoltage
	return fmt.Sprintf("%s.%d.InputRange.%d.%s", PrefixPsu, id, rangeId, InputRangesMinimumVoltage)
}

func (db *PsuFormat) GetInputRangesMaximumVoltageKey(id int64, rangeId int64) string {
	// Ex. Psu.0.InputRang.0.MaximumVoltage
	return fmt.Sprintf("%s.%d.InputRange.%d.%s", PrefixPsu, id, rangeId, InputRangesMaximumVoltage)
}

func (db *PsuFormat) GetInputRangesMinimumFrequencyHzKey(id int64, rangeId int64) string {
	// Ex. Psu.0.InputRang.0.MinimumFrequencyHz
	return fmt.Sprintf("%s.%d.InputRange.%d.%s", PrefixPsu, id, rangeId, InputRangesMinimumFrequencyHz)
}

func (db *PsuFormat) GetInputRangesMaximumFrequencyHzKey(id int64, rangeId int64) string {
	// Ex. Psu.0.InputRang.0.MaximumFrequencyHz
	return fmt.Sprintf("%s.%d.InputRange.%d.%s", PrefixPsu, id, rangeId, InputRangesMaximumFrequencyHz)
}

func (db *PsuFormat) GetInputRangesOutputWattageKey(id int64, rangeId int64) string {
	// Ex. Psu.0.InputRang.0.OutputWattage
	return fmt.Sprintf("%s.%d.InputRange.%d.%s", PrefixPsu, id, rangeId, InputRangesOutputWattage)
}
