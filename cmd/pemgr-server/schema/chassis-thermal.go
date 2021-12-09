package schema

import (
	"errors"
	"strconv"
)

type Thermal struct {
	// The type of a resource. [RO]
	OdataType string `json:"@odata.type"`

	// The identifier that uniquely identifies the Resource within
	// the collection of similar Resources. [RO]
	Id string `json:"Id"`

	// The name of the Resource or array member. [RO]
	Name string `json:"Name"`

	// The set of temperature sensors for this chassis. [RW]
	Temperatures []ThermalTemperatures

	// The set of fans for this chassis. [RW]
	Fans []ThermalFans

	// The OEM extension. [RW]
	// Oem ThermalOem `json:"Oem"`

	// The redundancy information for the set of fans in this chassis. [RW]
	Redundancy []ThermalRedundancy

	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`
}

type ThermalTemperatures struct {
	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`

	// The identifier for the member within the collection. [RO]
	MemberId string `json:"MemberId"`

	// The temperature sensor name. [RO]
	Name string `json:"Name,omitempty"`

	// The numerical identifier of the temperature sensor. [RO]
	SensorNumber int64 `json:"SensorNumber,omitempty"`

	// The status and health of a Resource and its children. [RW]
	Status CommonStatus `json:"Status"`

	// The temperature in degrees Celsius. [RO]
	ReadingCelsius float64 `json:"ReadingCelsius,omitempty"`

	// The value at which the reading is above normal range. [RO]
	UpperThresholdNonCritical float64 `json:"UpperThresholdNonCritical,omitempty"`

	// The value at which the reading is above normal range but not yet fatal. [RO]
	UpperThresholdCritical float64 `json:"UpperThresholdCritical,omitempty"`

	// The value at which the reading is above normal range and fatal. [RO]
	UpperThresholdFatal float64 `json:"UpperThresholdFatal,omitempty"`

	// The value at which the reading is below normal range. [RO]
	LowerThresholdNonCritical float64 `json:"LowerThresholdNonCritical,omitempty"`

	// The value at which the reading is below normal range but not yet fatal. [RO]
	LowerThresholdCritical float64 `json:"LowerThresholdCritical,omitempty"`

	// The value at which the reading is below normal range and fatal. [RO]
	LowerThresholdFatal float64 `json:"LowerThresholdFatal,omitempty"`

	// Minimum value for this sensor. [RO]
	MinReadingRangeTemp float64 `json:"MinReadingRangeTemp,omitempty"`

	// Maximum value for this sensor. [RO]
	MaxReadingRangeTemp float64 `json:"MaxReadingRangeTemp,omitempty"`

	// The area or device to which this temperature measurement applies. [RO]
	// Valid values:
	// ACInput:	                    An AC input.
	// ACMaintenanceBypassInput:	An AC maintenance bypass input.
	// ACOutput:	                An AC output.
	// ACStaticBypassInput:	        An AC static bypass input.
	// ACUtilityInput:	            An AC utility input.
	// ASIC:	                    An ASIC device, such as a networking chip or chipset component.
	// Accelerator:	                An accelerator.
	// Back:	                    The back of the chassis.
	// Backplane:	                A backplane within the chassis.
	// CPU:	                        A processor (CPU).
	// CPUSubsystem:	            The entire processor (CPU) subsystem.
	// Chassis:	                    The entire chassis.
	// ComputeBay:	                Within a compute bay.
	// CoolingSubsystem:	        The entire cooling, or air and liquid, subsystem.
	// DCBus:	                    A DC bus.
	// Exhaust:	                    The air exhaust point or points or region of the chassis.
	// ExpansionBay:	            Within an expansion bay.
	// FPGA:	                    An FPGA.
	// Fan:	                        A fan.
	// Front:	                    The front of the chassis.
	// GPU:	                        A graphics processor (GPU).
	// GPUSubsystem:	            The entire graphics processor (GPU) subsystem.
	// Intake:	                    The air intake point or points or region of the chassis.
	// LiquidInlet:	                The liquid inlet point of the chassis.
	// LiquidOutlet:	            The liquid outlet point of the chassis.
	// Lower:	                    The lower portion of the chassis.
	// Memory:	                    A memory device.
	// MemorySubsystem:	            The entire memory subsystem.
	// Motor:	                    A motor.
	// NetworkBay:	                Within a networking bay.
	// NetworkingDevice:	        A networking device.
	// PowerSubsystem:	            The entire power subsystem.
	// PowerSupply:	                A power supply.
	// PowerSupplyBay:	            Within a power supply bay.
	// Rectifier:	                A rectifier device.
	// Room:	                    The room.
	// StorageBay:	                Within a storage bay.
	// StorageDevice:	            A storage device.
	// SystemBoard:	                The system board (PCB).
	// Transformer:	                A transformer.
	// Upper:	                    The upper portion of the chassis.
	// VoltageRegulator:	        A voltage regulator device.
	PhysicalContext string `json:"PhysicalContext"`

	// The areas or devices to which this temperature applies. [RO]
	RelatedItem []CommonOid
}

type ThermalFans struct {
	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`

	// The identifier for the member within the collection. [RO]
	MemberId string `json:"MemberId"`

	// The temperature sensor name. [RO]
	Name string `json:"Name,omitempty"`

	// The numerical identifier of the temperature sensor. [RO]
	SensorNumber int64 `json:"SensorNumber,omitempty"`

	// The area or device to which this temperature measurement applies. [RO]
	// Valid values:
	// ACInput:	                    An AC input.
	// ACMaintenanceBypassInput:	An AC maintenance bypass input.
	// ACOutput:	                An AC output.
	// ACStaticBypassInput:	        An AC static bypass input.
	// ACUtilityInput:	            An AC utility input.
	// ASIC:	                    An ASIC device, such as a networking chip or chipset component.
	// Accelerator:	                An accelerator.
	// Back:	                    The back of the chassis.
	// Backplane:	                A backplane within the chassis.
	// CPU:	                        A processor (CPU).
	// CPUSubsystem:	            The entire processor (CPU) subsystem.
	// Chassis:	                    The entire chassis.
	// ComputeBay:	                Within a compute bay.
	// CoolingSubsystem:	        The entire cooling, or air and liquid, subsystem.
	// DCBus:	                    A DC bus.
	// Exhaust:	                    The air exhaust point or points or region of the chassis.
	// ExpansionBay:	            Within an expansion bay.
	// FPGA:	                    An FPGA.
	// Fan:	                        A fan.
	// Front:	                    The front of the chassis.
	// GPU:	                        A graphics processor (GPU).
	// GPUSubsystem:	            The entire graphics processor (GPU) subsystem.
	// Intake:	                    The air intake point or points or region of the chassis.
	// LiquidInlet:	                The liquid inlet point of the chassis.
	// LiquidOutlet:	            The liquid outlet point of the chassis.
	// Lower:	                    The lower portion of the chassis.
	// Memory:	                    A memory device.
	// MemorySubsystem:	            The entire memory subsystem.
	// Motor:	                    A motor.
	// NetworkBay:	                Within a networking bay.
	// NetworkingDevice:	        A networking device.
	// PowerSubsystem:	            The entire power subsystem.
	// PowerSupply:	                A power supply.
	// PowerSupplyBay:	            Within a power supply bay.
	// Rectifier:	                A rectifier device.
	// Room:	                    The room.
	// StorageBay:	                Within a storage bay.
	// StorageDevice:	            A storage device.
	// SystemBoard:	                The system board (PCB).
	// Transformer:	                A transformer.
	// Upper:	                    The upper portion of the chassis.
	// VoltageRegulator:	        A voltage regulator device.
	PhysicalContext string `json:"PhysicalContext"`

	// The status and health of a Resource and its children. [RW]
	Status CommonStatus `json:"Status"`

	// The fan speed. [RO]
	Reading int64 `json:"Reading"`

	// The units in which the fan reading and thresholds are measured. [RO]
	// Valid values:
	// Percent:	The fan reading and thresholds are measured as a percentage.
	// RPM:	    The fan reading and thresholds are measured in rotations per minute.
	ReadingUnits string `json:"ReadingUnits"`

	// The value at which the reading is below normal range and fatal. [RO]
	LowerThresholdFatal int64 `json:"LowerThresholdFatal"`

	// Minimum value for this sensor. [RO]
	MinReadingRange int64 `json:"MinReadingRange"`

	// Maximum value for this sensor. [RO]
	MaxReadingRange int64 `json:"MaxReadingRange"`

	// The set of redundancy groups for this fan. [RW]
	Redundancy []CommonOid

	// The areas or devices to which this temperature applies. [RO]
	RelatedItem []CommonOid

	Oem FanOem `json:"Oem"`
}

type ThermalLeds struct {
	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`

	// The identifier for the member within the collection. [RO]
	MemberId string `json:"MemberId"`

	// The temperature sensor name. [RO]
	Name string `json:"Name,omitempty"`

	// The numerical identifier of the temperature sensor. [RO]
	SensorNumber int64 `json:"SensorNumber,omitempty"`

	// The area or device to which this temperature measurement applies. [RO]
	// Valid values:
	// ACInput:	                    An AC input.
	// ACMaintenanceBypassInput:	An AC maintenance bypass input.
	// ACOutput:	                An AC output.
	// ACStaticBypassInput:	        An AC static bypass input.
	// ACUtilityInput:	            An AC utility input.
	// ASIC:	                    An ASIC device, such as a networking chip or chipset component.
	// Accelerator:	                An accelerator.
	// Back:	                    The back of the chassis.
	// Backplane:	                A backplane within the chassis.
	// CPU:	                        A processor (CPU).
	// CPUSubsystem:	            The entire processor (CPU) subsystem.
	// Chassis:	                    The entire chassis.
	// ComputeBay:	                Within a compute bay.
	// CoolingSubsystem:	        The entire cooling, or air and liquid, subsystem.
	// DCBus:	                    A DC bus.
	// Exhaust:	                    The air exhaust point or points or region of the chassis.
	// ExpansionBay:	            Within an expansion bay.
	// FPGA:	                    An FPGA.
	// Fan:	                        A fan.
	// Front:	                    The front of the chassis.
	// GPU:	                        A graphics processor (GPU).
	// GPUSubsystem:	            The entire graphics processor (GPU) subsystem.
	// Intake:	                    The air intake point or points or region of the chassis.
	// LiquidInlet:	                The liquid inlet point of the chassis.
	// LiquidOutlet:	            The liquid outlet point of the chassis.
	// Lower:	                    The lower portion of the chassis.
	// Memory:	                    A memory device.
	// MemorySubsystem:	            The entire memory subsystem.
	// Motor:	                    A motor.
	// NetworkBay:	                Within a networking bay.
	// NetworkingDevice:	        A networking device.
	// PowerSubsystem:	            The entire power subsystem.
	// PowerSupply:	                A power supply.
	// PowerSupplyBay:	            Within a power supply bay.
	// Rectifier:	                A rectifier device.
	// Room:	                    The room.
	// StorageBay:	                Within a storage bay.
	// StorageDevice:	            A storage device.
	// SystemBoard:	                The system board (PCB).
	// Transformer:	                A transformer.
	// Upper:	                    The upper portion of the chassis.
	// VoltageRegulator:	        A voltage regulator device.
	PhysicalContext string `json:"PhysicalContext"`

	// The status and health of a Resource and its children. [RW]
	Status CommonStatus `json:"Status"`

	// The fan speed. [RO]
	Reading int64 `json:"Reading"`

	// The units in which the fan reading and thresholds are measured. [RO]
	// Valid values:
	// Percent:	The fan reading and thresholds are measured as a percentage.
	// RPM:	    The fan reading and thresholds are measured in rotations per minute.
	ReadingUnits string `json:"ReadingUnits"`

	// The value at which the reading is below normal range and fatal. [RO]
	LowerThresholdFatal int64 `json:"LowerThresholdFatal"`

	// Minimum value for this sensor. [RO]
	MinReadingRange int64 `json:"MinReadingRange"`

	// Maximum value for this sensor. [RO]
	MaxReadingRange int64 `json:"MaxReadingRange"`

	// The set of redundancy groups for this fan. [RW]
	Redundancy []CommonOid

	// The areas or devices to which this temperature applies. [RO]
	RelatedItem []CommonOid
}

type ThermalButtons struct {
	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`

	// The identifier for the member within the collection. [RO]
	MemberId string `json:"MemberId"`

	// The temperature sensor name. [RO]
	Name string `json:"Name,omitempty"`

	// The numerical identifier of the temperature sensor. [RO]
	SensorNumber int64 `json:"SensorNumber,omitempty"`

	// The area or device to which this temperature measurement applies. [RO]
	// Valid values:
	// ACInput:	                    An AC input.
	// ACMaintenanceBypassInput:	An AC maintenance bypass input.
	// ACOutput:	                An AC output.
	// ACStaticBypassInput:	        An AC static bypass input.
	// ACUtilityInput:	            An AC utility input.
	// ASIC:	                    An ASIC device, such as a networking chip or chipset component.
	// Accelerator:	                An accelerator.
	// Back:	                    The back of the chassis.
	// Backplane:	                A backplane within the chassis.
	// CPU:	                        A processor (CPU).
	// CPUSubsystem:	            The entire processor (CPU) subsystem.
	// Chassis:	                    The entire chassis.
	// ComputeBay:	                Within a compute bay.
	// CoolingSubsystem:	        The entire cooling, or air and liquid, subsystem.
	// DCBus:	                    A DC bus.
	// Exhaust:	                    The air exhaust point or points or region of the chassis.
	// ExpansionBay:	            Within an expansion bay.
	// FPGA:	                    An FPGA.
	// Fan:	                        A fan.
	// Front:	                    The front of the chassis.
	// GPU:	                        A graphics processor (GPU).
	// GPUSubsystem:	            The entire graphics processor (GPU) subsystem.
	// Intake:	                    The air intake point or points or region of the chassis.
	// LiquidInlet:	                The liquid inlet point of the chassis.
	// LiquidOutlet:	            The liquid outlet point of the chassis.
	// Lower:	                    The lower portion of the chassis.
	// Memory:	                    A memory device.
	// MemorySubsystem:	            The entire memory subsystem.
	// Motor:	                    A motor.
	// NetworkBay:	                Within a networking bay.
	// NetworkingDevice:	        A networking device.
	// PowerSubsystem:	            The entire power subsystem.
	// PowerSupply:	                A power supply.
	// PowerSupplyBay:	            Within a power supply bay.
	// Rectifier:	                A rectifier device.
	// Room:	                    The room.
	// StorageBay:	                Within a storage bay.
	// StorageDevice:	            A storage device.
	// SystemBoard:	                The system board (PCB).
	// Transformer:	                A transformer.
	// Upper:	                    The upper portion of the chassis.
	// VoltageRegulator:	        A voltage regulator device.
	PhysicalContext string `json:"PhysicalContext"`

	// The status and health of a Resource and its children. [RW]
	Status CommonStatus `json:"Status"`

	// The fan speed. [RO]
	Reading int64 `json:"Reading"`

	// The units in which the fan reading and thresholds are measured. [RO]
	// Valid values:
	// Percent:	The fan reading and thresholds are measured as a percentage.
	// RPM:	    The fan reading and thresholds are measured in rotations per minute.
	ReadingUnits string `json:"ReadingUnits"`

	// The value at which the reading is below normal range and fatal. [RO]
	LowerThresholdFatal int64 `json:"LowerThresholdFatal"`

	// Minimum value for this sensor. [RO]
	MinReadingRange int64 `json:"MinReadingRange"`

	// Maximum value for this sensor. [RO]
	MaxReadingRange int64 `json:"MaxReadingRange"`

	// The set of redundancy groups for this fan. [RW]
	Redundancy []CommonOid

	// The areas or devices to which this temperature applies. [RO]
	RelatedItem []CommonOid
}

type ThermalRedundancy struct {
	// The unique identifier for a resource. [RO]
	OdataId string `json:"@odata.id"`

	// The identifier for the member within the collection. [RO]
	MemberId string `json:"MemberId"`

	// The temperature sensor name. [RO]
	Name string `json:"Name,omitempty"`

	// The links to components of this redundancy set. [RO]
	RedundancySet []CommonOid

	// The redundancy mode of the group. [RW]
	// Valid values:
	// Failover:	    Failure of one unit automatically causes a standby or offline unit in the redundancy set to take over its functions.
	// N+m:	            Multiple units are available and active such that normal operation will continue if one or more units fail.
	// NotRedundant:	The subsystem is not configured in a redundancy mode,
	//                  either due to configuration or the functionality has been disabled by the user.
	// Sharing:	        Multiple units contribute or share such that operation will continue, but at a reduced capacity, if one or more units fail.
	// Sparing:	        One or more spare units are available to take over the function of a failed unit, but takeover is not automatic.
	Mode string `json:"Mode"`

	// The status and health of a Resource and its children. [RW]
	Status CommonStatus `json:"Status"`

	// The minumum number of members needed for this group to be redundant. [RO]
	MinNumNeeded int64 `json:"MinNumNeeded,omitempty"`

	// The maximum number of members allowable for this particular redundancy group. [RO]
	MaxNumSupported int64 `json:"MaxNumSupported,omitempty"`
}

type FanOem struct {
	Custom FanOemCustom `json:"Custom"`
}

type FanOemCustom struct {
	// The duty cycle of the fan. [RW]
	// Valid values:
	// "10"-"100":	The duty cycle can be set from 10 to 100 (Unit:%).
	Duty string `json:"Duty"`
}

func (fan FanOemCustom) Validation() error {
	if len(fan.Duty) == 0 {
		return errors.New("The duty is empty")
	}

	duty, err := strconv.ParseInt(fan.Duty, 10, 32)
	if err != nil {
		return errors.New("The duty is not a number")
	}

	if duty < 0 || duty > 100 {
		return errors.New("The duty is not a valid number")
	}
	return nil
}
